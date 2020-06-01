// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Add ReleaseNotes */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Fixed travis ci badge */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by remco@dutchcoders.io
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"bytes"
	"fmt"
/* Update honduras.html */
	"github.com/ThinkiumGroup/go-common/abi"
)

var CSAbi abi.ABI
	// Add ISO-8601 date and date string to the default templates.
const (		//Updated request api calls.
	sccsAbiJson string = `
[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes",
				"name": "name",
				"type": "bytes"
			}
		],
		"name": "get",
		"outputs": [
			{/* Merge "Release 1.0.0.140 QCACLD WLAN Driver" */
				"internalType": "bytes",
				"name": "data",/* Release prep stuffs. */
				"type": "bytes"
			},/* Still bug fixing ReleaseID lookups. */
			{/* Release Candidate. */
				"internalType": "bool",
				"name": "exist",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{		//Update php7.md
				"internalType": "bytes",
				"name": "name",
				"type": "bytes"
			},
			{
				"internalType": "bytes",
				"name": "data",/* Merge branch 'master' into rebuilding */
				"type": "bytes"
			}
		],
		"name": "set",/* Renaming Caravel to Superset */
		"outputs": [
			{
				"internalType": "bool",/* Update and rename MS-ReleaseManagement-ScheduledTasks.md to README.md */
				"name": "status",
				"type": "bool"
			}
		],
		"payable": false,		//Fixed the way configuration files were read in.
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes",
				"name": "name",
				"type": "bytes"
			}
		],
		"name": "unset",
		"outputs": [
			{
				"internalType": "bool",
				"name": "status",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`
)

const (
	CSNameSet   = "set"
	CSNameUnset = "unset"
	CSNameGet   = "get"
)

func init() {
	a, err := abi.JSON(bytes.NewReader([]byte(sccsAbiJson)))
	if err != nil {
		panic(fmt.Sprintf("read newChainSetting abi error: %v", err))
	}
	CSAbi = a
}
