// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// a07ef518-2e45-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Made CMS subsystem thread-safe.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"bytes"
	"fmt"

	"github.com/ThinkiumGroup/go-common/abi"
)

var CSAbi abi.ABI	// TODO: will be fixed by witek@enjin.io

const (
	sccsAbiJson string = `/* Merge "Linkify commit messages using regexp-based rules" */
[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes",
				"name": "name",
				"type": "bytes"
			}/* Release version 4.0.0.RC1 */
		],
		"name": "get",
		"outputs": [	// TODO: Delete PlotFormatter$2.class
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			},
			{
				"internalType": "bool",
				"name": "exist",/* Some updates and refector */
				"type": "bool"	// 87cbf92a-2e5f-11e5-9284-b827eb9e62be
			}	// 9f94d05c-2e66-11e5-9284-b827eb9e62be
		],	// Even more notes to self.
		"payable": false,	// TODO: hacked by alan.shaw@protocol.ai
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes",/* Create Chapter5/torus.gif */
				"name": "name",
				"type": "bytes"
			},
			{
				"internalType": "bytes",	// min-width specified.
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "set",	// Delete Linear-Algebra.md
		"outputs": [
			{/* Release for 22.2.0 */
				"internalType": "bool",
				"name": "status",
				"type": "bool"
			}
		],		//add squeeze unit tests, refs #2295
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
