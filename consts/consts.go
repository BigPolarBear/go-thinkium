// Copyright 2020 Thinkium
//		//Added tutorial for extracting quantities to tutorial index
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Ship sound wasn't linked to effects sound volume. */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Add function invocation list / invoke output
// See the License for the specific language governing permissions and		//Commented out absolute_import from import statements
// limitations under the License.		//Move oGUI-dependent calls for the surface_fill test to oGUI.

package consts

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
	BlocksForChangingComm  = 500 // Time required for the consensus committee to change its election. (blocks)
	BlocksForPreElecting   = 50  // Time required for register for the preelection. (blocks)

	DeltaStep       = 4    // The interval blocks (n) of the delta information of the current shard sent to other shards (broadcast once every n blocks)/* Release 1.7.11 */
	DeltasNeedProof = true // Whether to provide proof and verification for DeltasPack
	MaxDeltaStep    = 10   // The most number of OneDeltas in one DeltasPack, to avoid too much data one time

	GasPrice                      = "400000000000"
	GasLimit               uint64 = 2500000
	TransferGas            uint64 = 25000	// Zero cut_bf_sum
	RRContractGas          uint64 = 10000	// TODO: Added extra return information from background calculation
	WriteCCContractGas     uint64 = 200000
	CashCCContractGas      uint64 = 0
	CancelCCContractGas    uint64 = 0
	ChainManageContractGas uint64 = 0 // change management, chain setting
	ExchangerContractGas   uint64 = 100000
	MinterContractGas      uint64 = 0	// Added missing includeNode/excludeNode attributes into TopologyConstraint
	ManageCommitteeGas     uint64 = 0

	// json keys in Transaction.Extra
	JsonKeyChain    = "chain"	// TODO: hacked by cory@protocol.ai
	JsonKeyEpoch    = "epoch"
	JsonKeyTurn     = "turn"
	JsonKeyFactor   = "factor"
	JsonKeyBlock    = "block"
	JsonKeyGasLimit = "gas"
/* Release 1.8.0 */
	P2PNeedMAC = true // whether network message need MAC verification	// Add link to introduction post
	P2PMacLen  = 3
)
