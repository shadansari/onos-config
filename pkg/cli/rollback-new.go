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

package cli

import (
	"context"
	"github.com/onosproject/onos-config/pkg/northbound/admin"
	"github.com/spf13/cobra"
)

func getRollbackNewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollback-new <changeId>",
		Short: "Rolls-back a new network change",
		Args:  cobra.MaximumNArgs(1),
		RunE:  runNewRollbackCommand,
	}
	return cmd
}

func runNewRollbackCommand(cmd *cobra.Command, args []string) error {
	clientConnection, clientConnectionError := getConnection()

	if clientConnectionError != nil {
		return clientConnectionError
	}
	client := admin.CreateConfigAdminServiceClient(clientConnection)

	changeID := ""
	if len(args) == 1 {
		changeID = args[0]
	}

	resp, err := client.RollbackNewNetworkChange(
		context.Background(), &admin.RollbackRequest{Name: changeID})
	if err != nil {
		return err
	}
	Output("Rollback success %s\n", resp.Message)
	return nil
}