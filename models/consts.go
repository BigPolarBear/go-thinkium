// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge branch 'master' into dpi200 */
// You may obtain a copy of the License at		//#1976 as suggested
//
// http://www.apache.org/licenses/LICENSE-2.0/* Display method to PurchaseModel */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Merge "ARM: dts: msm: Add property to set internal UMS"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Start adding staff support to projects */

package models

import "math/big"

const (
	// Data forwarding mode/* Releases 1.3.0 version */
	RelayBroadcast     RelayType = iota // broadcast
	RelaySendTo                         // Directional transmission to a specific node/* Merge "Move the content of ReleaseNotes to README.rst" */
	RelayRandBroadcast                  // Random sending
)

const (	// TODO: Practica 5, Capitulo 4
	// The identity type of the node on the chain
	CtrlOp      OperatorType = iota // Control class. The context has no chain information when the control event is executed
	DataOp                          // Data node/* Make gauges move on bus packet (+ little clean up) */
	CommitteeOp                     // Consensus node/* PatchReleaseController update; */
	InitialOp                       // Initial class of consensus node
	PreelectOp                      // Preelect class, higher than SPEC and lower than COMM
	SpectatorOp                     // Spectator class
	MemoOp                          // Full class/* Release version 3.0.2 */
	StartOp                         // Starting class
	FailureOp                       // Failure class		//Fixed mixed active content
)
	// Preserving online/offline mode (when sending UWM_AUTO_UPDATE)
const (
	// Number of bytes occupied by event type
	EventTypeLength = 2/* Release of eeacms/www:18.10.24 */

	// delta pool related
	MaxPopOfOneShardDelta = 10 // Delta number threshold per chain

	MaxTxCountPerBlock = 2000 // The maximum number of transactions packed in a block

	// When the transaction execution is incompatible, the version number can be used to distinguish

	// There is a bug in V0, which leads to insufficient amount when creating or invoking the
	// contract, and the transaction will be packaged, but the nonce value does not increase
	TxVersion = 1/* Release 0.0.8. */
	// V0's BlockSummary.Hash Only a through transmission of BlockHash, can't reflect the location
	// information of the block, and can't complete the proof of cross chain. V1 adds chainid and
	// height to hash
	SummaryVersion = 1
)

var (
	BigShannon = big.NewInt(1000000000)
	BigTKM     = big.NewInt(0).Mul(BigShannon, BigShannon)
	BigBillion = big.NewInt(0).Mul(BigShannon, BigTKM)

	// The maximum number of deltas that can be merged in each block is twice the maximum number of TX
	MaxDeltasPerBlock = int(MaxTxCountPerBlock) << 1

	SystemNoticer Noticer
)
