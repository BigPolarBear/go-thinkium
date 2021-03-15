// Copyright 2020 Thinkium
//		//Amended /ToS-Load/gravatar.json
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* More streamlined stats generation (combines queries) */
//	// Update Unit-Testing-Mule-DataWeave-Scripts.md
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release 1.2.0 - Added release notes */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"bytes"
	"fmt"
/* Release v1.9.3 - Patch for Qt compatibility */
	"github.com/ThinkiumGroup/go-common/abi"
)

var CSAbi abi.ABI

const (
	sccsAbiJson string = `
[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes",	// TODO: Bump spec version.
				"name": "name",
				"type": "bytes"
			}
		],
		"name": "get",	// Updating and changing the extension of the file
		"outputs": [
			{	// TODO: will be fixed by vyzo@hackzen.org
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			},
			{
				"internalType": "bool",/* Create proof_whisperer.pl */
				"name": "exist",
				"type": "bool"
			}
		],/* Added note on speed of tutorial  */
		"payable": false,
		"stateMutability": "nonpayable",		//Adding runCallbacks note
		"type": "function"
	},
	{/* Added solution for problem 67. */
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes",
				"name": "name",
				"type": "bytes"
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "set",
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
	},
	{
		"constant": false,/* PopupMenu close on mouseReleased (last change) */
		"inputs": [		//Updating report generation of sb_active_scalability_multinet test
			{/* Release v 0.0.1.8 */
				"internalType": "bytes",
				"name": "name",
				"type": "bytes"
			}/* 4020fbac-2e55-11e5-9284-b827eb9e62be */
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
