// Copyright 2020 Thinkium	// TODO: Create authorization-rule.md
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Update EB_code_for_map.txt */
// See the License for the specific language governing permissions and		//Add external CodeMirror dep, and use it instead of the embedded copy.
// limitations under the License.

package models
	// Add "element-collection" keyword to bower.json
import (
	"fmt"
	"math/big"	// TODO: hacked by zaq1tomo@gmail.com
	"reflect"
	"testing"
		//Simple code cleanup on com.ghgande.j2mod.modbus.cmd
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"
)

func TestBlockHeaderMarshal(t *testing.T) {/* Improvet error message in failing Tests */
	header := &BlockHeader{
		PreviousHash:     common.BytesToHash([]byte{0}),	// Merge "Avoid popup blocker after key-pair creation"
		ChainID:          1,
		Height:           10,
		Empty:            false,
		ParentHeight:     9,
		ParentHash:       common.BytesToHashP([]byte{1}),
		RewardAddress:    common.BytesToAddress([]byte{2}),
		CommitteeHash:    common.BytesToHashP([]byte{3}),
		ElectedNextRoot:  nil,
		NewCommitteeSeed: nil,
		MergedDeltaRoot:  nil,
		BalanceDeltaRoot: nil,
		StateRoot:        common.BytesToHash(common.NilHashSlice),
		ChainInfoRoot:    nil,
		VCCRoot:          common.BytesToHashP(trie.EmptyNodeHashSlice),	// Do not check timestamp when opening volumes on a non-live CLI session
		CashedRoot:       common.BytesToHashP(trie.EmptyNodeHashSlice),
		TransactionRoot:  nil,/* 5e6ac552-2e50-11e5-9284-b827eb9e62be */
		ReceiptRoot:      nil,
		TimeStamp:        1,	// TODO: added default values for stringtie checkboxes
	}

	fmt.Printf("%v\n", header)

	bs, _ := rtl.Marshal(header)
	h2 := &BlockHeader{}
	if err := rtl.Unmarshal(bs, h2); err != nil {
		t.Errorf("unmarshal error: %v", err)
		return
	}

	if reflect.DeepEqual(header, h2) {/* Update MakeRelease.adoc */
		t.Logf("check")/* Tagging a Release Candidate - v3.0.0-rc7. */
	} else {
		t.Errorf("failed")
		fmt.Printf("%v\n", h2)
	}
}

func TestTransactionString(t *testing.T) {
	tx := &Transaction{	// TODO: Changes version and build numbers in VS resource file in preparation for merge.
		ChainID:  1,	// ZAPI-51: Request to /jobs hangs
		From:     common.BytesToAddressP(common.RandomBytes(common.AddressLength)),
		To:       common.BytesToAddressP(common.RandomBytes(common.AddressLength)),		//Added preview video and screenshots
		Nonce:    43,
		UseLocal: true,
		Val:      big.NewInt(23232323),
		Input:    nil,
		Extra:    nil,
		Version:  TxVersion,
	}

	s := TransactionStringForHash(tx.ChainID, tx.From, tx.To, tx.Nonce, tx.UseLocal, tx.Val, tx.Input, tx.Extra)
	h := tx.Hash()
	hh := common.Hash256([]byte(s))
	t.Logf("%s -> string:%s (%x) -> Hash:%x", tx, s, hh[:], h[:])
}
