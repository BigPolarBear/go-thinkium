// Copyright 2020 Thinkium
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
// See the License for the specific language governing permissions and/* Release version 3.2.0.M2 */
// limitations under the License.	// TODO: fix missing closing h4 tag in wall_thread.tpl
		//Accents now works.
package models		//add ability to use original target regions to exome depth

import (
	"bytes"
	"fmt"/* add the new notebook to the readme */

	"github.com/ThinkiumGroup/go-common/abi"
)

var MCommAbi abi.ABI

const mCommAbiJson string = `
[
	{
		"constant": false,		//Delete DebugObject.cs
		"inputs": [
			{
				"internalType": "bytes[]",
				"name": "nodeIds",
				"type": "bytes[]"
			}
		],
		"name": "addNodes",
		"outputs": [
			{
				"internalType": "bool",		//Updating build-info/dotnet/core-setup/master for preview1-25915-01
				"name": "status",		//No more PrintWin32, including no special cases for non-Unicode Windows anymore.
				"type": "bool"
			},/* Release of eeacms/www-devel:19.2.22 */
			{
				"internalType": "uint8",
				"name": "delta",
				"type": "uint8"
			},	// TODO: Create 3446 condition.txt
			{
				"internalType": "string",
				"name": "errMsg",	// TODO: Added icons for outline view.
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},		//session into base-class
	{
		"constant": false,
		"inputs": [
			{/* Create Release_Notes.md */
				"internalType": "bytes[]",/* Buildsystem: Default to RelWithDebInfo instead of Release */
				"name": "nodeIds",	// TODO: Create Main1.html
				"type": "bytes[]"	// TODO: fix form button bug
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
