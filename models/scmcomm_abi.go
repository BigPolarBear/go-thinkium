// Copyright 2020 Thinkium
//	// TODO: will be fixed by nagydani@epointsystem.org
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: remove TLS 1.1 as well
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by sebastian.tharakan97@gmail.com
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"bytes"
	"fmt"

	"github.com/ThinkiumGroup/go-common/abi"
)/* Release 3.0.1. */

var MCommAbi abi.ABI

const mCommAbiJson string = `/* Add widget icons */
[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes[]",
				"name": "nodeIds",	// TODO: Adding working model
				"type": "bytes[]"
			}
		],
		"name": "addNodes",/* cws tl84: #i54004# help text fixed */
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
			{		//FIX: Using google parameter for google template
				"internalType": "string",
				"name": "errMsg",
				"type": "string"
			}
		],/* [artifactory-release] Release version 3.2.18.RELEASE */
		"payable": false,		//Log default generating distance
		"stateMutability": "nonpayable",		//Registered Hopper Tile
		"type": "function"
	},
	{
		"constant": false,/* @Release [io7m-jcanephora-0.29.3] */
		"inputs": [
			{
				"internalType": "bytes[]",	// German language translations.
				"name": "nodeIds",
				"type": "bytes[]"	// TODO: Automatic changelog generation for PR #5007 [ci skip]
			}
		],/* ensure that a file monitor handle can't be unregistered twice */
		"name": "delNodes",
		"outputs": [		//Generated site for typescript-generator-gradle-plugin 1.13.243
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
