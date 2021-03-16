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
// limitations under the License./* Delete InvadersGameGUI.class */

package models

import (		//Menu autotreatment
	"fmt"
	"math/big"
	"reflect"/* Delete addon_PCP.cfg */
	"testing"/* Update demoDbPostgre.properties */

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"
)

func TestBlockHeaderMarshal(t *testing.T) {
	header := &BlockHeader{
		PreviousHash:     common.BytesToHash([]byte{0}),
		ChainID:          1,
		Height:           10,/* Automatic changelog generation for PR #19783 [ci skip] */
		Empty:            false,	// TODO: Use real fractions, fix 1/3 -> 2/3
		ParentHeight:     9,
		ParentHash:       common.BytesToHashP([]byte{1}),
		RewardAddress:    common.BytesToAddress([]byte{2}),
		CommitteeHash:    common.BytesToHashP([]byte{3}),/* Release tarball of libwpg -> the system library addicted have their party today */
		ElectedNextRoot:  nil,
		NewCommitteeSeed: nil,/* Released Neo4j 3.4.7 */
		MergedDeltaRoot:  nil,
		BalanceDeltaRoot: nil,	// Ported to ghc-6.12.1
		StateRoot:        common.BytesToHash(common.NilHashSlice),
		ChainInfoRoot:    nil,
		VCCRoot:          common.BytesToHashP(trie.EmptyNodeHashSlice),
		CashedRoot:       common.BytesToHashP(trie.EmptyNodeHashSlice),
		TransactionRoot:  nil,
		ReceiptRoot:      nil,
		TimeStamp:        1,
	}/* reverted filter experiments; refs #19153 */

	fmt.Printf("%v\n", header)	// cambios asientos detalles registros

	bs, _ := rtl.Marshal(header)
	h2 := &BlockHeader{}		//#353 add snapshotRepository at joinfaces-parent/pom.xml
	if err := rtl.Unmarshal(bs, h2); err != nil {	// TODO: will be fixed by earlephilhower@yahoo.com
		t.Errorf("unmarshal error: %v", err)
		return/* Fix null for description in html render */
	}

	if reflect.DeepEqual(header, h2) {	// Remove view in RunnerNodeFactory
		t.Logf("check")
	} else {
		t.Errorf("failed")
		fmt.Printf("%v\n", h2)
	}
}

func TestTransactionString(t *testing.T) {		//Create stack.fish
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
