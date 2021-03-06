// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onosproject/onos-topo/api/topo (interfaces: TopoClient,Topo_ListClient,Topo_SubscribeClient,TopoServer,Topo_ListServer,Topo_SubscribeServer)

// Package device is a generated GoMock package.
package device

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	topo "github.com/onosproject/onos-topo/api/topo"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockTopoClient is a mock of TopoClient interface
type MockTopoClient struct {
	ctrl     *gomock.Controller
	recorder *MockTopoClientMockRecorder
}

// MockTopoClientMockRecorder is the mock recorder for MockTopoClient
type MockTopoClientMockRecorder struct {
	mock *MockTopoClient
}

// NewMockTopoClient creates a new mock instance
func NewMockTopoClient(ctrl *gomock.Controller) *MockTopoClient {
	mock := &MockTopoClient{ctrl: ctrl}
	mock.recorder = &MockTopoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopoClient) EXPECT() *MockTopoClientMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockTopoClient) Delete(arg0 context.Context, arg1 *topo.DeleteRequest, arg2 ...grpc.CallOption) (*topo.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*topo.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockTopoClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTopoClient)(nil).Delete), varargs...)
}

// Get mocks base method
func (m *MockTopoClient) Get(arg0 context.Context, arg1 *topo.GetRequest, arg2 ...grpc.CallOption) (*topo.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*topo.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTopoClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTopoClient)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockTopoClient) List(arg0 context.Context, arg1 *topo.ListRequest, arg2 ...grpc.CallOption) (topo.Topo_ListClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(topo.Topo_ListClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockTopoClientMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTopoClient)(nil).List), varargs...)
}

// Set mocks base method
func (m *MockTopoClient) Set(arg0 context.Context, arg1 *topo.SetRequest, arg2 ...grpc.CallOption) (*topo.SetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(*topo.SetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockTopoClientMockRecorder) Set(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockTopoClient)(nil).Set), varargs...)
}

// Subscribe mocks base method
func (m *MockTopoClient) Subscribe(arg0 context.Context, arg1 *topo.SubscribeRequest, arg2 ...grpc.CallOption) (topo.Topo_SubscribeClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(topo.Topo_SubscribeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockTopoClientMockRecorder) Subscribe(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockTopoClient)(nil).Subscribe), varargs...)
}

// MockTopo_ListClient is a mock of Topo_ListClient interface
type MockTopo_ListClient struct {
	ctrl     *gomock.Controller
	recorder *MockTopo_ListClientMockRecorder
}

// MockTopo_ListClientMockRecorder is the mock recorder for MockTopo_ListClient
type MockTopo_ListClientMockRecorder struct {
	mock *MockTopo_ListClient
}

// NewMockTopo_ListClient creates a new mock instance
func NewMockTopo_ListClient(ctrl *gomock.Controller) *MockTopo_ListClient {
	mock := &MockTopo_ListClient{ctrl: ctrl}
	mock.recorder = &MockTopo_ListClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopo_ListClient) EXPECT() *MockTopo_ListClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockTopo_ListClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockTopo_ListClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTopo_ListClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockTopo_ListClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTopo_ListClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTopo_ListClient)(nil).Context))
}

// Header mocks base method
func (m *MockTopo_ListClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockTopo_ListClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTopo_ListClient)(nil).Header))
}

// Recv mocks base method
func (m *MockTopo_ListClient) Recv() (*topo.ListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*topo.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockTopo_ListClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTopo_ListClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockTopo_ListClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTopo_ListClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTopo_ListClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockTopo_ListClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTopo_ListClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTopo_ListClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockTopo_ListClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockTopo_ListClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockTopo_ListClient)(nil).Trailer))
}

// MockTopo_SubscribeClient is a mock of Topo_SubscribeClient interface
type MockTopo_SubscribeClient struct {
	ctrl     *gomock.Controller
	recorder *MockTopo_SubscribeClientMockRecorder
}

// MockTopo_SubscribeClientMockRecorder is the mock recorder for MockTopo_SubscribeClient
type MockTopo_SubscribeClientMockRecorder struct {
	mock *MockTopo_SubscribeClient
}

// NewMockTopo_SubscribeClient creates a new mock instance
func NewMockTopo_SubscribeClient(ctrl *gomock.Controller) *MockTopo_SubscribeClient {
	mock := &MockTopo_SubscribeClient{ctrl: ctrl}
	mock.recorder = &MockTopo_SubscribeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopo_SubscribeClient) EXPECT() *MockTopo_SubscribeClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockTopo_SubscribeClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockTopo_SubscribeClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTopo_SubscribeClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockTopo_SubscribeClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTopo_SubscribeClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTopo_SubscribeClient)(nil).Context))
}

// Header mocks base method
func (m *MockTopo_SubscribeClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockTopo_SubscribeClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTopo_SubscribeClient)(nil).Header))
}

// Recv mocks base method
func (m *MockTopo_SubscribeClient) Recv() (*topo.SubscribeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*topo.SubscribeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockTopo_SubscribeClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTopo_SubscribeClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockTopo_SubscribeClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTopo_SubscribeClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTopo_SubscribeClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockTopo_SubscribeClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTopo_SubscribeClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTopo_SubscribeClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockTopo_SubscribeClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockTopo_SubscribeClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockTopo_SubscribeClient)(nil).Trailer))
}

// MockTopoServer is a mock of TopoServer interface
type MockTopoServer struct {
	ctrl     *gomock.Controller
	recorder *MockTopoServerMockRecorder
}

// MockTopoServerMockRecorder is the mock recorder for MockTopoServer
type MockTopoServerMockRecorder struct {
	mock *MockTopoServer
}

// NewMockTopoServer creates a new mock instance
func NewMockTopoServer(ctrl *gomock.Controller) *MockTopoServer {
	mock := &MockTopoServer{ctrl: ctrl}
	mock.recorder = &MockTopoServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopoServer) EXPECT() *MockTopoServerMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockTopoServer) Delete(arg0 context.Context, arg1 *topo.DeleteRequest) (*topo.DeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*topo.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockTopoServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTopoServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockTopoServer) Get(arg0 context.Context, arg1 *topo.GetRequest) (*topo.GetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*topo.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTopoServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTopoServer)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockTopoServer) List(arg0 *topo.ListRequest, arg1 topo.Topo_ListServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockTopoServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTopoServer)(nil).List), arg0, arg1)
}

// Set mocks base method
func (m *MockTopoServer) Set(arg0 context.Context, arg1 *topo.SetRequest) (*topo.SetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(*topo.SetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockTopoServerMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockTopoServer)(nil).Set), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockTopoServer) Subscribe(arg0 *topo.SubscribeRequest, arg1 topo.Topo_SubscribeServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockTopoServerMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockTopoServer)(nil).Subscribe), arg0, arg1)
}

// MockTopo_ListServer is a mock of Topo_ListServer interface
type MockTopo_ListServer struct {
	ctrl     *gomock.Controller
	recorder *MockTopo_ListServerMockRecorder
}

// MockTopo_ListServerMockRecorder is the mock recorder for MockTopo_ListServer
type MockTopo_ListServerMockRecorder struct {
	mock *MockTopo_ListServer
}

// NewMockTopo_ListServer creates a new mock instance
func NewMockTopo_ListServer(ctrl *gomock.Controller) *MockTopo_ListServer {
	mock := &MockTopo_ListServer{ctrl: ctrl}
	mock.recorder = &MockTopo_ListServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopo_ListServer) EXPECT() *MockTopo_ListServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockTopo_ListServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTopo_ListServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTopo_ListServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockTopo_ListServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTopo_ListServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTopo_ListServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockTopo_ListServer) Send(arg0 *topo.ListResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockTopo_ListServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTopo_ListServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockTopo_ListServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockTopo_ListServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockTopo_ListServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockTopo_ListServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTopo_ListServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTopo_ListServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockTopo_ListServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockTopo_ListServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockTopo_ListServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockTopo_ListServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockTopo_ListServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockTopo_ListServer)(nil).SetTrailer), arg0)
}

// MockTopo_SubscribeServer is a mock of Topo_SubscribeServer interface
type MockTopo_SubscribeServer struct {
	ctrl     *gomock.Controller
	recorder *MockTopo_SubscribeServerMockRecorder
}

// MockTopo_SubscribeServerMockRecorder is the mock recorder for MockTopo_SubscribeServer
type MockTopo_SubscribeServerMockRecorder struct {
	mock *MockTopo_SubscribeServer
}

// NewMockTopo_SubscribeServer creates a new mock instance
func NewMockTopo_SubscribeServer(ctrl *gomock.Controller) *MockTopo_SubscribeServer {
	mock := &MockTopo_SubscribeServer{ctrl: ctrl}
	mock.recorder = &MockTopo_SubscribeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopo_SubscribeServer) EXPECT() *MockTopo_SubscribeServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockTopo_SubscribeServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTopo_SubscribeServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTopo_SubscribeServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockTopo_SubscribeServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTopo_SubscribeServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTopo_SubscribeServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockTopo_SubscribeServer) Send(arg0 *topo.SubscribeResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockTopo_SubscribeServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTopo_SubscribeServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockTopo_SubscribeServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockTopo_SubscribeServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockTopo_SubscribeServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockTopo_SubscribeServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTopo_SubscribeServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTopo_SubscribeServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockTopo_SubscribeServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockTopo_SubscribeServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockTopo_SubscribeServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockTopo_SubscribeServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockTopo_SubscribeServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockTopo_SubscribeServer)(nil).SetTrailer), arg0)
}
