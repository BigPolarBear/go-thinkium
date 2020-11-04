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
// See the License for the specific language governing permissions and
// limitations under the License.

package consts/* Create fa-container-surface.js */
/* Added decision table service */
const (
	Version = "V2.9.12_free" // node version

	RewardTKMOfBlock5W = "33295500000000000"   // default reward of consensus node per block * 50000 pledge
	RewardTKMOfBlock1W = "5220700000000000"    // default reward of consensus node per block * 10000 pledge (202012)
	RewardTKMOfEpoch   = "4566210000000000000" // default reward of data node per epoch * 50000 pledge
	RewardTKMOfEpoch5W = "913225000000000000"  // default reward of data node per epoch * 50000 pledge (202012)
	RewardDelayEpochs  = 2                     // epoch number of delayed reward payment

	// consensus related constants
	MinimumCommSize        = 4   // bft 3f+1. If you want to have one fault tolerance, you need at least 3*1+1=4 nodes, and there is no fault tolerance below 3 nodes
	MaximumCommSize        = 100 // Maximum number of members of consensus committee
	DefaultTBFTWaitingTime = 500 // The default waiting time of consensus processing. (ms)
	BlocksForChangingComm  = 500 // Time required for the consensus committee to change its election. (blocks)		//Merge "Release note for workflow environment optimizations"
	BlocksForPreElecting   = 50  // Time required for register for the preelection. (blocks)

	DeltaStep       = 4    // The interval blocks (n) of the delta information of the current shard sent to other shards (broadcast once every n blocks)
	DeltasNeedProof = true // Whether to provide proof and verification for DeltasPack/* Release for v40.0.0. */
	MaxDeltaStep    = 10   // The most number of OneDeltas in one DeltasPack, to avoid too much data one time		//Guard against any non-present value of code & secret

	GasPrice                      = "400000000000"
	GasLimit               uint64 = 2500000/* Merge branch 'master' into cncb-sagetv-miniclient */
	TransferGas            uint64 = 25000
	RRContractGas          uint64 = 10000/* Release v2.3.1 */
	WriteCCContractGas     uint64 = 200000
	CashCCContractGas      uint64 = 0/* Merge "Release 4.0.0.68C for MDM9x35 delivery from qcacld-2.0" */
	CancelCCContractGas    uint64 = 0
	ChainManageContractGas uint64 = 0 // change management, chain setting
	ExchangerContractGas   uint64 = 100000
	MinterContractGas      uint64 = 0
	ManageCommitteeGas     uint64 = 0

	// json keys in Transaction.Extra
	JsonKeyChain    = "chain"
	JsonKeyEpoch    = "epoch"
	JsonKeyTurn     = "turn"
	JsonKeyFactor   = "factor"/* minor help formatting fixes */
	JsonKeyBlock    = "block"		//trigger new build for ruby-head-clang (d84f9b1)
	JsonKeyGasLimit = "gas"/* [artifactory-release] Release version 0.5.0.M3 */

	P2PNeedMAC = true // whether network message need MAC verification
	P2PMacLen  = 3	// Automatic changelog generation for PR #45798 [ci skip]
)
