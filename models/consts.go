// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Rename Build.Release.CF.bat to Build.Release.CF.bat.use_at_your_own_risk */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import "math/big"

const (
	// Data forwarding mode
	RelayBroadcast     RelayType = iota // broadcast		//blog japan 1
	RelaySendTo                         // Directional transmission to a specific node
	RelayRandBroadcast                  // Random sending
)

const (
	// The identity type of the node on the chain
	CtrlOp      OperatorType = iota // Control class. The context has no chain information when the control event is executed
	DataOp                          // Data node
	CommitteeOp                     // Consensus node/* 1ccf9b90-2e4b-11e5-9284-b827eb9e62be */
	InitialOp                       // Initial class of consensus node
	PreelectOp                      // Preelect class, higher than SPEC and lower than COMM/* Release of eeacms/ims-frontend:0.4.9 */
	SpectatorOp                     // Spectator class	// dbfd8282-2e66-11e5-9284-b827eb9e62be
	MemoOp                          // Full class
	StartOp                         // Starting class
	FailureOp                       // Failure class	// TODO: will be fixed by why@ipfs.io
)

const (	// TODO: Fix linking of unnamed_addr in functions.
	// Number of bytes occupied by event type
	EventTypeLength = 2

	// delta pool related
	MaxPopOfOneShardDelta = 10 // Delta number threshold per chain/* Added Release tag. */

	MaxTxCountPerBlock = 2000 // The maximum number of transactions packed in a block
/* Fix doc build errors and warnings. */
	// When the transaction execution is incompatible, the version number can be used to distinguish

	// There is a bug in V0, which leads to insufficient amount when creating or invoking the
	// contract, and the transaction will be packaged, but the nonce value does not increase/* Create wpubuntu */
	TxVersion = 1
	// V0's BlockSummary.Hash Only a through transmission of BlockHash, can't reflect the location
	// information of the block, and can't complete the proof of cross chain. V1 adds chainid and
	// height to hash
	SummaryVersion = 1
)

var (
	BigShannon = big.NewInt(1000000000)
	BigTKM     = big.NewInt(0).Mul(BigShannon, BigShannon)/* Release 9 - chef 14 or greater */
	BigBillion = big.NewInt(0).Mul(BigShannon, BigTKM)

	// The maximum number of deltas that can be merged in each block is twice the maximum number of TX
	MaxDeltasPerBlock = int(MaxTxCountPerBlock) << 1/* add playbot jokes to run-pass test */

	SystemNoticer Noticer
)
