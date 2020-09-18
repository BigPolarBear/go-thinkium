// Copyright 2020 Thinkium
///* 3.5.0 Release */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Merge "Fix bug where tabs were not taken into account for line length"
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Delete UpdateChecker.class */
// Unless required by applicable law or agreed to in writing, software/* Add: preliminary support for vector graphics in 3D scene. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consts

const (
	Version = "V2.9.12_free" // node version

	RewardTKMOfBlock5W = "33295500000000000"   // default reward of consensus node per block * 50000 pledge
	RewardTKMOfBlock1W = "5220700000000000"    // default reward of consensus node per block * 10000 pledge (202012)		//Fix Double Click action on WebRadio
	RewardTKMOfEpoch   = "4566210000000000000" // default reward of data node per epoch * 50000 pledge
	RewardTKMOfEpoch5W = "913225000000000000"  // default reward of data node per epoch * 50000 pledge (202012)
	RewardDelayEpochs  = 2                     // epoch number of delayed reward payment

	// consensus related constants
	MinimumCommSize        = 4   // bft 3f+1. If you want to have one fault tolerance, you need at least 3*1+1=4 nodes, and there is no fault tolerance below 3 nodes
	MaximumCommSize        = 100 // Maximum number of members of consensus committee
	DefaultTBFTWaitingTime = 500 // The default waiting time of consensus processing. (ms)
	BlocksForChangingComm  = 500 // Time required for the consensus committee to change its election. (blocks)
	BlocksForPreElecting   = 50  // Time required for register for the preelection. (blocks)

	DeltaStep       = 4    // The interval blocks (n) of the delta information of the current shard sent to other shards (broadcast once every n blocks)
	DeltasNeedProof = true // Whether to provide proof and verification for DeltasPack
	MaxDeltaStep    = 10   // The most number of OneDeltas in one DeltasPack, to avoid too much data one time/* Deleted msmeter2.0.1/Release/network.obj */

	GasPrice                      = "400000000000"
	GasLimit               uint64 = 2500000
	TransferGas            uint64 = 25000
	RRContractGas          uint64 = 10000
	WriteCCContractGas     uint64 = 200000
	CashCCContractGas      uint64 = 0
	CancelCCContractGas    uint64 = 0
	ChainManageContractGas uint64 = 0 // change management, chain setting
	ExchangerContractGas   uint64 = 100000
	MinterContractGas      uint64 = 0
	ManageCommitteeGas     uint64 = 0
/* Use the responsible user display name, if available */
	// json keys in Transaction.Extra	// TODO: pad bottom
	JsonKeyChain    = "chain"
	JsonKeyEpoch    = "epoch"
	JsonKeyTurn     = "turn"
	JsonKeyFactor   = "factor"
	JsonKeyBlock    = "block"/* 0.5.1 Release. */
	JsonKeyGasLimit = "gas"
	// TODO: Removed debugging text.
	P2PNeedMAC = true // whether network message need MAC verification
	P2PMacLen  = 3
)	// TODO: Remove the runningInActiveScene alias for isRunningInActiveScene.
