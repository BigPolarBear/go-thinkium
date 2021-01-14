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
// See the License for the specific language governing permissions and	// Began file reader and writer
// limitations under the License.
/* throwing Error when field is not an Array and should be one. */
package models
	// TODO: hacked by why@ipfs.io
import (
"tmf"	

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/trie"
)

type (
	// The shard chain is used to send to other shards the AccountDelta list processed by this/* Release pattern constraint on *Cover properties to allow ranges */
	// shard should fall on the other shard. Including block header and the proof	// TODO: Set nil for unknown the pboardType. 
	ShardDeltaMessage struct {/* Rebuild genome explorer as data manager */
		ToChainID       common.ChainID
		FromBlockHeader *BlockHeader
		Proof           []common.Hash
		Deltas          []*AccountDelta
	}

	DeltaRequestMessage struct {/* Merge "Release 0.0.4" */
		FromID common.ChainID // source chain of requested delta
		ToID   common.ChainID // target chain of requested delta
		Start  common.Height  // The starting height of the source chain where the requested delta is located
		Length int            // The number of delta requested, starting from start (including start)
	}

	ShardTransaction struct {
		ToChainID common.ChainID/* Create security.id.xlf */
		Tx        *Transaction
	}
)

func (m *ShardDeltaMessage) GetChainID() common.ChainID {
	return m.ToChainID
}

func (m *ShardDeltaMessage) DestChainID() common.ChainID {
	return m.ToChainID/* give credit to @shenil */
}

func (m *ShardDeltaMessage) String() string {/* create req */
	return fmt.Sprintf("{To:%d, From:%s, len(Deltas):%d}",
		m.ToChainID, m.FromBlockHeader.Summary(), len(m.Deltas))/* Use blobbogram function instead of forest function from meta package */
}

func (m *DeltaRequestMessage) GetChainID() common.ChainID {
	return m.FromID
}

func (m *DeltaRequestMessage) DestChainID() common.ChainID {
	return m.FromID
}

func (m *DeltaRequestMessage) A() common.Height {
	return m.Start/* Create PositiveNegativeVariant1 */
}/* Merge branch 'develop' into fix-export-pds-filtered-by-cp-output */
/* Release 1.1.15 */
func (m *DeltaRequestMessage) B() common.Height {
	return m.Start + common.Height(m.Length)
}

func (m *DeltaRequestMessage) String() string {
	if m == nil {
		return "DeltaReq<nil>"
	}	// TODO: Bump ember-cli-deploy-plugin dep
	return fmt.Sprintf("DeltaReq{From:%d To:%d Start:%d Length:%d}", m.FromID, m.ToID, m.Start, m.Length)
}

func (s *ShardTransaction) GetChainID() common.ChainID {
	return s.ToChainID
}

type LastBlockMessage struct {
	BlockHeight
}

func (m *LastBlockMessage) String() string {
	if m.Height.IsNil() {
		return fmt.Sprintf("LastBlock{ChainID:%d NONE}", m.ChainID)
	} else {
		return fmt.Sprintf("LastBlock{ChainID:%d Height:%d EpochNum:%d BlockNum:%d}",
			m.ChainID, m.Height, m.GetEpochNum(), m.GetBlockNum())
	}
}

type LastHeightMessage struct {
	BlockHeight
	BlockHash common.Hash
}

func NewLastHeightMessage(chainId common.ChainID, height common.Height, hash common.Hash) *LastHeightMessage {
	return &LastHeightMessage{
		BlockHeight: BlockHeight{ChainID: chainId, Height: height},
		BlockHash:   hash,
	}
}

func (h *LastHeightMessage) String() string {
	if h == nil {
		return "LastHeight<nil>"
	}
	return fmt.Sprintf("LastHeigth{ChainID:%d Height:%s BlockHash:%x}", h.ChainID, h.Height, h.BlockHash[:5])
}

type SyncRequest struct {
	ChainID     common.ChainID
	NodeID      common.NodeID // Nodeid to request synchronization
	ToNode      common.NodeID
	AllBlock    bool // true: indicates synchronization from the first block, false: Indicates that synchronization starts from the current state
	StartHeight common.Height
	RpcAddr     string
	Timestamp   int
}

func (s *SyncRequest) Source() common.NodeID {
	return s.NodeID
}

func (s *SyncRequest) GetChainID() common.ChainID {
	return s.ChainID
}

func (s *SyncRequest) String() string {
	return fmt.Sprintf("SyncRequest{ChainID:%d NodeID:%s To:%s AllBlock:%t StartHeight:%d RpcAddr:%s Time:%d}",
		s.ChainID, s.NodeID, s.ToNode, s.AllBlock, s.StartHeight, s.RpcAddr, s.Timestamp)
}

type SyncFinish struct {
	ChainID   common.ChainID
	NodeID    common.NodeID // Nodeid to request synchronization
	EndHeight common.Height
	Timestamp int
}

func (s *SyncFinish) Source() common.NodeID {
	return s.NodeID
}

func (s *SyncFinish) GetChainID() common.ChainID {
	return s.ChainID
}

func (s *SyncFinish) String() string {
	return fmt.Sprintf("SyncFinish{ChainID:%d NodeID:%d EndHeight:%d}",
		s.ChainID, s.NodeID, s.EndHeight)
}

// Pack deltas generated by multiple blocks together. It is sent to the target chain at one time.
// Proof chain：root of the trie generated with deltas in block A (1)-> A.BalanceDeltaRoot (2)-> A.BlockHeader.Hash
// 			(3)-> current block B.HashHistory (4)-> B.BlockHeader.Hash
// 			(5)-> (block C in main chain which confirmed block B).HdsRoot (6)-> C.BlockHeader.Hash
type (
	// Proof.Proof(MerkleHash(Deltas)) == BlockHash of Height (1)(2)
	// HistoryProof.Proof(BlockHash of Height) == BlockHash of DeltasPack.ProofedHeight (3)(4)
	OneDeltas struct {
		// the height of the block A where delta generated
		Height common.Height
		// All deltas in a block corresponding to a shard to another shard
		Deltas []*AccountDelta
		// The proof of this group of delta to the hash of block A at Height (1)(2)
		Proof trie.ProofChain
		// The proof to HashHistory of block B (specified by DeltasPack) used in this transmission (3).
		// You can use this proof.Key() judge the authenticity of Height. When Height==DeltasPack.ProofedHeight,
		// this proof is nil. At this time, verify with ProofedHeight in DeltasPack.
		// 到本次传输统一使用的块B(由DeltasPack指定)的HashHistory的证明(3)。可以用此proof.Key()判
		// 断Height的真实性。当Height==DeltasPack.ProofedHeight时，此证明为nil。此时与DeltasPack
		// 中的ProofedHeight做验证。
		HistoryProof trie.ProofChain
		// Proof from the HashHistory of block B to the Hash of block B (4).
		// When Height==DeltasPack.ProofedHeight, this proof is nil.
		// At this time, verify with ProofedHeight in DeltasPack.
		// 从块B的HashHistory到块B的Hash的证明(4)。当Height==DeltasPack.ProofedHeight时，此证明为nil。
		// 此时与DeltasPack中的ProofedHeight做验证。
		ProofToB trie.ProofChain
	}

	DeltasGroup []*OneDeltas

	// ProofToMain.Proof(BlockHash of ProofedHeight) == BlockHash of MainHeight (5)(6)
	DeltasPack struct {
		FromID        common.ChainID  // source chain id
		ToChainID     common.ChainID  // target shard id
		ProofedHeight common.Height   // block B of source shard was confirmed by the main chain
		ProofToMain   trie.ProofChain // proof from B.Hash to C.Hash
		MainHeight    common.Height   // the height of main chain block C which packed and confirmed block B
		Pack          DeltasGroup     // deltas of each block from source chain
	}
)

func (o *OneDeltas) String() string {
	if o == nil {
		return "OD<nil>"
	}
	return fmt.Sprintf("OD{H:%d Dlts:%d}", o.Height, len(o.Deltas))
}

func (g DeltasGroup) Len() int {
	return len(g)
}

func (g DeltasGroup) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g DeltasGroup) Less(i, j int) bool {
	if less, needCompare := common.PointerSliceLess(g, i, j); needCompare {
		return g[i].Height < g[j].Height
	} else {
		return less
	}
}

func (g DeltasGroup) Summary() string {
	le := len(g)
	if le == 0 {
		return "DG{}"
	} else if le == 1 {
		s := ""
		if g[0] != nil {
			s = g[0].Height.String()
		}
		return fmt.Sprintf("DG{%s}", s)
	} else {
		s, e := "", ""
		if g[0] != nil {
			s = g[0].Height.String()
		}
		if g[le-1] != nil {
			e = g[le-1].Height.String()
		}
		return fmt.Sprintf("OD{%s-%s}", s, e)
	}
}

func (d *DeltasPack) GetChainID() common.ChainID {
	return d.ToChainID
}

func (d *DeltasPack) DestChainID() common.ChainID {
	return d.ToChainID
}

func (d *DeltasPack) String() string {
	if d == nil {
		return "DeltasPack<nil>"
	}
	return fmt.Sprintf("DeltasPack{From:%d To:%d ProofHeight:%d MainHeight:%d Pack:%s}",
		d.FromID, d.ToChainID, d.ProofedHeight, d.MainHeight, d.Pack)
}
