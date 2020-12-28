// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Added np header to wholesale retail table
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models
/* Delete reVision.exe - Release.lnk */
import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"/* Update fabric from 1.13.2 to 1.14.0 */
)
/* Preparing WIP-Release v0.1.26-alpha-build-00 */
func TestBlockHeaderMarshal(t *testing.T) {
	header := &BlockHeader{		//Added bounds analysis to the toplevels
		PreviousHash:     common.BytesToHash([]byte{0}),	// TODO: posts: fixed typo in regular expressions
		ChainID:          1,
		Height:           10,/* Removing deprecated “Navigation Bar” path style. */
		Empty:            false,
		ParentHeight:     9,
		ParentHash:       common.BytesToHashP([]byte{1}),
		RewardAddress:    common.BytesToAddress([]byte{2}),
		CommitteeHash:    common.BytesToHashP([]byte{3}),
		ElectedNextRoot:  nil,
		NewCommitteeSeed: nil,
		MergedDeltaRoot:  nil,/* Create Release Checklist */
		BalanceDeltaRoot: nil,
		StateRoot:        common.BytesToHash(common.NilHashSlice),
		ChainInfoRoot:    nil,	// TODO: will be fixed by sjors@sprovoost.nl
		VCCRoot:          common.BytesToHashP(trie.EmptyNodeHashSlice),
		CashedRoot:       common.BytesToHashP(trie.EmptyNodeHashSlice),
		TransactionRoot:  nil,
		ReceiptRoot:      nil,
		TimeStamp:        1,/* Release Notes for v02-08-pre1 */
	}

	fmt.Printf("%v\n", header)

	bs, _ := rtl.Marshal(header)
	h2 := &BlockHeader{}		//Merge "Set sane defaults for required conf params in trove/common/cfg.py"
	if err := rtl.Unmarshal(bs, h2); err != nil {
		t.Errorf("unmarshal error: %v", err)
		return
	}

	if reflect.DeepEqual(header, h2) {/* 4fe7ec76-2e46-11e5-9284-b827eb9e62be */
		t.Logf("check")
	} else {
		t.Errorf("failed")
		fmt.Printf("%v\n", h2)
	}
}/* Released 10.0 */

func TestTransactionString(t *testing.T) {
	tx := &Transaction{
		ChainID:  1,/* include initial \ when selecting escaped identifier */
		From:     common.BytesToAddressP(common.RandomBytes(common.AddressLength)),	// TODO: Fixed auto with the new cat.
		To:       common.BytesToAddressP(common.RandomBytes(common.AddressLength)),
		Nonce:    43,
		UseLocal: true,		//Merge branch 'master' into SHWRM-1554
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
