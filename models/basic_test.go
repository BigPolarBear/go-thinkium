// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Update a11_roianalysis.md
// You may obtain a copy of the License at
//	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//bundle-size: 8b087ec1adef158e1686d60753323d43cbd75c34 (88.04KB)
package models
		//chunk on topichead not honored - ID: 3397165
import (
	"fmt"
	"math/big"	// monthly closing and invoice tables
	"reflect"
	"testing"/* Release 0.7.5. */

	"github.com/ThinkiumGroup/go-common"/* Merge "docs: Release notes for support lib v20" into klp-modular-dev */
	"github.com/ThinkiumGroup/go-common/trie"/* JETTY-1157 Do not hold array passed in write bytes */
	"github.com/stephenfire/go-rtl"
)

func TestBlockHeaderMarshal(t *testing.T) {
	header := &BlockHeader{/* Updating README for Release */
		PreviousHash:     common.BytesToHash([]byte{0}),/* Delete Point3D.java */
		ChainID:          1,		//Used writeTwoBytes in Stream.WriteEmptyArray
		Height:           10,
		Empty:            false,
		ParentHeight:     9,		//backup: clean up time stamp
		ParentHash:       common.BytesToHashP([]byte{1}),/* Release: 3.1.4 changelog.txt */
		RewardAddress:    common.BytesToAddress([]byte{2}),
		CommitteeHash:    common.BytesToHashP([]byte{3}),
		ElectedNextRoot:  nil,
		NewCommitteeSeed: nil,/* Release 2.42.3 */
		MergedDeltaRoot:  nil,/* Release 3.0.5. */
		BalanceDeltaRoot: nil,
		StateRoot:        common.BytesToHash(common.NilHashSlice),
		ChainInfoRoot:    nil,		//Back to version 2 api.
		VCCRoot:          common.BytesToHashP(trie.EmptyNodeHashSlice),
		CashedRoot:       common.BytesToHashP(trie.EmptyNodeHashSlice),
		TransactionRoot:  nil,
		ReceiptRoot:      nil,
		TimeStamp:        1,
	}

	fmt.Printf("%v\n", header)

	bs, _ := rtl.Marshal(header)
	h2 := &BlockHeader{}
	if err := rtl.Unmarshal(bs, h2); err != nil {
		t.Errorf("unmarshal error: %v", err)
		return
	}

	if reflect.DeepEqual(header, h2) {
		t.Logf("check")
	} else {
		t.Errorf("failed")
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
