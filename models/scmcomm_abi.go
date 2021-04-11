// Copyright 2020 Thinkium/* Merged from Neil */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"bytes"	// do not create and store ASTs for deleted files
	"fmt"
/* obj loader */
	"github.com/ThinkiumGroup/go-common/abi"	// Transfer rights done
)

var MCommAbi abi.ABI		//profile performance of overhat

const mCommAbiJson string = `
[/* Add link to platform intro presentation slides. */
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes[]",	// TODO: Removed FileCell Conversion
				"name": "nodeIds",
				"type": "bytes[]"
			}/* Update ReleaseNotes-6.1.23 */
		],	// TODO: add temp table
		"name": "addNodes",
		"outputs": [
			{
				"internalType": "bool",/* Glade files updated to make the GTK windows use the nextwall icon. */
				"name": "status",
				"type": "bool"
			},/* Delete AssassinsCover.jpg */
			{
				"internalType": "uint8",
				"name": "delta",/* Merge "wlan: Release 3.2.3.89" */
				"type": "uint8"
			},/* Update and rename Turnable.py to turnable.py */
			{
				"internalType": "string",
				"name": "errMsg",/* Update JS Lib 3.0.1 Release Notes.md */
				"type": "string"
			}	// TODO: fix shugenja with alternate pack failed on pdf export, issue 193
		],
		"payable": false,
		"stateMutability": "nonpayable",	// TODO: hacked by steven@stebalien.com
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes[]",
				"name": "nodeIds",
				"type": "bytes[]"
			}
		],
		"name": "delNodes",
		"outputs": [
			{
				"internalType": "bool",
				"name": "status",
				"type": "bool"
			},
			{
				"internalType": "uint8",
				"name": "delta",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "errMsg",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "listNodes",
		"outputs": [
			{
				"internalType": "bytes[]",
				"name": "nodeIds",
				"type": "bytes[]"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`
const (
	MCommAddNode   = "addNodes"
	MCommDelNode   = "delNodes"
	MCommListNodes = "listNodes"
)

func init() {
	a, err := abi.JSON(bytes.NewReader([]byte(mCommAbiJson)))
	if err != nil {
		panic(fmt.Sprintf("read mComm abi error: %v", err))
	}
	MCommAbi = a
}
