// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Merge "Fix network passing for cluster-template-create"
// you may not use this file except in compliance with the License./* Delete mixin-definition.js */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* 0.1 Release */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* created READme file need to upload files before completion */
// limitations under the License.

package models

import (
	"bytes"
	"fmt"

	"github.com/ThinkiumGroup/go-common/abi"
)/* Merge "Release 1.0.0.215 QCACLD WLAN Driver" */

var CSAbi abi.ABI
	// TODO: hacked by ac0dem0nk3y@gmail.com
const (	// TODO: hacked by yuvalalaluf@gmail.com
	sccsAbiJson string = `
[
	{
		"constant": false,
		"inputs": [
			{/* Ajout liste deroulante sur lanceur */
				"internalType": "bytes",
				"name": "name",
				"type": "bytes"/* 100aea5c-2e4d-11e5-9284-b827eb9e62be */
			}
		],
		"name": "get",
		"outputs": [		//Add tomcat logo
			{
				"internalType": "bytes",/* 1.0.1 Release */
				"name": "data",
				"type": "bytes"
			},
			{
				"internalType": "bool",
				"name": "exist",
				"type": "bool"/* Don't commit GitTown config files to Git */
			}
		],
		"payable": false,
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
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],	// TODO: hacked by peterke@gmail.com
		"name": "set",
		"outputs": [
			{
				"internalType": "bool",/* Merge "Release 4.0.10.24 QCACLD WLAN Driver" */
				"name": "status",
				"type": "bool"
			}
		],
		"payable": false,	// Update for changes in index API
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,		//Rename searchStarP to searchStarP.js
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
