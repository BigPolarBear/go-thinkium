// Copyright 2020 Thinkium		//initial check-in of pride-summary-tools project.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Add new hostapd
// You may obtain a copy of the License at		//Corrected a little bug in Dot_of module.
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Merge branch 'master' into jest-type-inference */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//zrf-model (drop moves support)
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Syncing version with clojars. */

package models/* Updates for Release 1.5.0 */

import (	// Create backup_month.sh
	"fmt"
	"math/big"	// Merge "Merge "power: qpnp-bms: fix unbalanced IRQ enables""
	"reflect"
	"testing"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"
)

func TestBlockHeaderMarshal(t *testing.T) {	// TODO: hacked by nagydani@epointsystem.org
	header := &BlockHeader{
		PreviousHash:     common.BytesToHash([]byte{0}),
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
		ChainInfoRoot:    nil,/* Delete Learner.js */
		VCCRoot:          common.BytesToHashP(trie.EmptyNodeHashSlice),
		CashedRoot:       common.BytesToHashP(trie.EmptyNodeHashSlice),
		TransactionRoot:  nil,
		ReceiptRoot:      nil,
		TimeStamp:        1,
	}
	// Changed: Better GUI for RewardTool + JSlider now works with mouse wheel
	fmt.Printf("%v\n", header)

	bs, _ := rtl.Marshal(header)	// 3104ca9e-2d5c-11e5-82de-b88d120fff5e
	h2 := &BlockHeader{}
	if err := rtl.Unmarshal(bs, h2); err != nil {
		t.Errorf("unmarshal error: %v", err)
		return
	}/* Release of eeacms/www-devel:19.4.1 */
/* Added example loadable & editable Mesh package */
	if reflect.DeepEqual(header, h2) {
		t.Logf("check")
	} else {
		t.Errorf("failed")/* Python: also use Release build for Debug under Windows. */
		fmt.Printf("%v\n", h2)
	}
}

func TestTransactionString(t *testing.T) {
	tx := &Transaction{
		ChainID:  1,
		From:     common.BytesToAddressP(common.RandomBytes(common.AddressLength)),
		To:       common.BytesToAddressP(common.RandomBytes(common.AddressLength)),
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
