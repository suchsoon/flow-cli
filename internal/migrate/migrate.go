/*
 * Flow CLI
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package migrate

import (
	"github.com/onflow/flow-go-sdk"
	"github.com/spf13/cobra"
)

func init() {
	getStagedCodeCommand.AddToParent(Cmd)
	IsStagedCommand.AddToParent(Cmd)
	listStagedContractsCommand.AddToParent(Cmd)
	stageContractCommand.AddToParent(Cmd)
	unstageContractCommand.AddToParent(Cmd)
}

var Cmd = &cobra.Command{
	Use:              "migrate",
	Short:            "Migrate your Cadence project to 1.0",
	TraverseChildren: true,
	GroupID:          "migrate",
}

// address of the migration contract on each network
var migrationContractStagingAddress = map[string]string{
	"testnet":   "0x2ceae959ed1a7e7a",
	"crescendo": "0x27b2302520211b67",
}

// MigrationContractStagingAddress returns the address of the migration contract on the given network
func MigrationContractStagingAddress(network string) flow.Address {
	return flow.HexToAddress(migrationContractStagingAddress[network])
}
