// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package synchronizer

import (
	"context"
	"fmt"
	"math"
	"strconv"
	syncPrimitives "sync"
	"time"

	devicechange "github.com/onosproject/onos-config/api/types/change/device"
	devicetype "github.com/onosproject/onos-config/api/types/device"
	"github.com/onosproject/onos-config/pkg/dispatcher"
	"github.com/onosproject/onos-config/pkg/events"
	"github.com/onosproject/onos-config/pkg/modelregistry"
	"github.com/onosproject/onos-config/pkg/southbound"
	"github.com/onosproject/onos-config/pkg/store/change/device"
	"github.com/onosproject/onos-config/pkg/utils"
	topodevice "github.com/onosproject/onos-topo/api/device"
	"github.com/onosproject/onos-topo/api/topo"
)

var connectionMonitors = make(map[topo.ID]*connectionMonitor)
var connections = make(map[topo.ID]bool)

const backoffInterval = 10 * time.Millisecond
const maxBackoffTime = 5 * time.Second

// Factory is a go routine thread that listens out for Device creation
// and deletion events and spawns Synchronizer threads for them
// These connectionMonitors then listen out for configEvents relative to a device and
func Factory(topoChannel <-chan *topo.SubscribeResponse, opStateChan chan<- events.OperationalStateEvent,
	southboundErrorChan chan<- events.DeviceResponse, dispatcher *dispatcher.Dispatcher,
	modelRegistry *modelregistry.ModelRegistry, operationalStateCache map[topo.ID]devicechange.TypedValueMap,
	newTargetFn func() southbound.TargetIf,
	operationalStateCacheLock *syncPrimitives.RWMutex, deviceChangeStore device.Store) {

	errChan := make(chan events.DeviceResponse)
	for {
		select {
		case topoEvent, ok := <-topoChannel:
			if !ok {
				return
			}

			device := topo.ObjectToDevice(topoEvent.Update.Object)
			connMon, ok := connectionMonitors[topo.ID(device.ID)]
			if !ok && topoEvent.Update.Type != topo.Update_DELETE {
				log.Infof("Topo device %s %s", device.ID, topoEvent.Update.Type)
				connMon = &connectionMonitor{
					opStateChan:               opStateChan,
					southboundErrorChan:       errChan,
					dispatcher:                dispatcher,
					modelRegistry:             modelRegistry,
					operationalStateCache:     operationalStateCache,
					operationalStateCacheLock: operationalStateCacheLock,
					deviceChangeStore:         deviceChangeStore,
					device:                    device,
					target:                    newTargetFn(),
				}
				connectionMonitors[topo.ID(device.ID)] = connMon
				go connMon.connect()
			} else if ok && topoEvent.Update.Type == topo.Update_MODIFY {
				changed := false
				if topoEvent.Update.Object.Attributes["address"] != "" && connMon.device.Address != topoEvent.Update.Object.Attributes["address"] {
					oldAddress := connMon.device.Address
					connMon.device.Address = topoEvent.Update.Object.Attributes["address"]
					changed = true
					log.Infof("Topo device %s is being UPDATED - waiting to complete", device.ID)
					connMon.close()
					// TODO Change grpc.DialContext() used to non blocking so that we can
					//  close the connection right away See https://github.com/onosproject/onos-config/issues/981
					waitTime := *connMon.device.GetTimeout() //Use the old timeout in case it has changed
					if maxBackoffTime > waitTime {
						waitTime = maxBackoffTime
					}
					time.Sleep(waitTime + time.Millisecond*20) // close might not take effect until timeout
					go connMon.reconnect()
					log.Infof("Topo device %s UPDATED address %s -> %s ", device.ID, oldAddress, topoEvent.Update.Object.Attributes["address"])
				}
				if topoEvent.Update.Object.Attributes["timeout"] != "" && connMon.device.Timeout.String() != topoEvent.Update.Object.Attributes["timeout"] {
					connMon.mu.Lock()
					oldTimeout := connMon.device.Timeout
					t, _ := strconv.Atoi(topoEvent.Update.Object.Attributes["timeout"])
					duration := time.Duration(t) * time.Second
					connMon.device.Timeout = &duration
					changed = true
					connMon.mu.Unlock()
					log.Infof("Topo device %s UPDATED timeout %s -> %s ", device.ID, oldTimeout, topoEvent.Update.Object.Attributes["timeout"])
				}
				if len(connMon.device.Protocols) != len(topoEvent.Update.Object.GetEntity().Protocols) {
					// Ignoring any topo protocol updates - we set the gNMI one and
					// Don't really care about the others
					changed = true
				}
				if !changed {
					log.Infof("Topo device %s UPDATE not supported %v", device.ID, device)
					southboundErrorChan <- events.NewErrorEventNoChangeID(
						events.EventTypeTopoUpdate, string(device.ID),
						fmt.Errorf("topo update event ignored %v", topoEvent))
				}
			} else if ok && topoEvent.Update.Type == topo.Update_DELETE {
				log.Infof("Topo device %s is being REMOVED - waiting to complete", device.ID)
				delete(connectionMonitors, topo.ID(device.ID))
				delete(connections, topo.ID(device.ID))
				connMon.close()
				// TODO Change grpc.DialContext() used to non blocking so that we can
				//  close the connection right away See https://github.com/onosproject/onos-config/issues/981
				waitTime := time.Duration(math.Max(float64(*connMon.device.GetTimeout()), float64(maxBackoffTime)))
				time.Sleep(waitTime + 100*time.Millisecond)
				log.Infof("Topo device %s REMOVED after %s", device.ID, waitTime)
			} else {
				log.Warnf("Unhandled event from topo service %v", topoEvent)
			}
		case event, ok := <-errChan:
			if !ok {
				return
			}

			log.Infof("Received event %v", event)
			deviceID := topo.ID(event.Subject())
			switch event.EventType() {
			case events.EventTypeErrorDeviceConnect:
				deviceID := topo.ID(event.Subject())
				connMon, ok := connectionMonitors[deviceID]
				if ok && connections[deviceID] {
					connections[deviceID] = false
					go connMon.reconnect()
				}
			case events.EventTypeDeviceConnected:
				connections[deviceID] = true
			}
			southboundErrorChan <- event
		}
	}
}

// connectionMonitor reacts to device events to establish connections to the device
type connectionMonitor struct {
	opStateChan               chan<- events.OperationalStateEvent
	southboundErrorChan       chan<- events.DeviceResponse
	dispatcher                *dispatcher.Dispatcher
	modelRegistry             *modelregistry.ModelRegistry
	operationalStateCache     map[topo.ID]devicechange.TypedValueMap
	operationalStateCacheLock *syncPrimitives.RWMutex
	deviceChangeStore         device.Store
	device                    *topodevice.Device
	target                    southbound.TargetIf
	cancel                    context.CancelFunc
	closed                    bool
	mu                        syncPrimitives.RWMutex
}

func (cm *connectionMonitor) reconnect() {
	cm.mu.Lock()
	if cm.cancel != nil {
		cm.cancel()
		cm.cancel = nil
	}
	cm.mu.Unlock()
	cm.operationalStateCacheLock.Lock()
	delete(cm.operationalStateCache, topo.ID(cm.device.ID))
	cm.operationalStateCacheLock.Unlock()
	cm.connect()
}

func (cm *connectionMonitor) connect() {
	count := 0
	for {
		count++

		cm.mu.Lock()
		closed := cm.closed
		cm.mu.Unlock()

		if closed {
			return
		}

		err := cm.synchronize()
		if err != nil {
			backoffTime := time.Duration(math.Min(float64(backoffInterval)*math.Pow(2, float64(count)), float64(maxBackoffTime)))
			log.Infof("Failed to connect to %s. Retry after %v Attempt %d", cm.device.ID, backoffTime, count)
			time.Sleep(backoffTime)
		} else {
			return
		}
	}
}

// synchronize connects to the device for synchronization
func (cm *connectionMonitor) synchronize() error {
	ctx, cancel := context.WithCancel(context.Background())
	cm.mu.Lock()
	cm.cancel = cancel
	cm.mu.Unlock()

	cm.mu.RLock()
	log.Infof("Connecting to device %v", cm.device)
	modelName := utils.ToModelName(devicetype.Type(cm.device.Type), devicetype.Version(cm.device.Version))
	mReadOnlyPaths, ok := cm.modelRegistry.ModelReadOnlyPaths[modelName]
	if !ok {
		log.Warnf("Cannot check for read only paths for target %cm with %cm because "+
			"Model Plugin not available - continuing", cm.device.ID, cm.device.Version)
	}
	mStateGetMode := modelregistry.GetStateOpState // default
	mPlugin, ok := cm.modelRegistry.ModelPlugins[modelName]
	if !ok {
		log.Warnf("Cannot check for StateGetMode for target %cm with %cm because "+
			"Model Plugin not available - continuing", cm.device.ID, cm.device.Version)
	} else {
		mStateGetMode = modelregistry.GetStateMode(mPlugin.GetStateMode())
	}
	valueMap := make(devicechange.TypedValueMap)
	cm.operationalStateCacheLock.Lock()
	cm.operationalStateCache[topo.ID(cm.device.ID)] = valueMap
	cm.operationalStateCacheLock.Unlock()
	cm.mu.RUnlock()

	sync, err := New(ctx, cm.device, cm.opStateChan, cm.southboundErrorChan,
		valueMap, mReadOnlyPaths, cm.target, mStateGetMode, cm.operationalStateCacheLock, cm.deviceChangeStore)
	if err != nil {
		log.Errorf("Error connecting to device %v: %v", cm.device, err)
		//unregistering the listener for changes to the device
		//unregistering the listener for changes to the device
		cm.dispatcher.UnregisterOperationalState(string(cm.device.ID))
		cm.operationalStateCacheLock.Lock()
		delete(cm.operationalStateCache, topo.ID(cm.device.ID))
		cm.operationalStateCacheLock.Unlock()
		return err
	}

	//spawning two go routines to propagate changes and to get operational state
	//go sync.syncConfigEventsToDevice(target, respChan)
	cm.southboundErrorChan <- events.NewDeviceConnectedEvent(events.EventTypeDeviceConnected, string(cm.device.ID))
	if sync.getStateMode == modelregistry.GetStateOpState {
		go sync.syncOperationalStateByPartition(ctx, cm.target, cm.southboundErrorChan)
	} else if sync.getStateMode == modelregistry.GetStateExplicitRoPaths ||
		sync.getStateMode == modelregistry.GetStateExplicitRoPathsExpandWildcards {
		go sync.syncOperationalStateByPaths(ctx, cm.target, cm.southboundErrorChan)
	}
	return nil
}

// close closes the synchronizer
func (cm *connectionMonitor) close() {
	cm.mu.Lock()
	cm.closed = true
	if cm.cancel != nil {
		cm.cancel()
		cm.cancel = nil
	}
	cm.mu.Unlock()
	cm.operationalStateCacheLock.Lock()
	delete(cm.operationalStateCache, topo.ID(cm.device.ID))
	cm.operationalStateCacheLock.Unlock()
}
