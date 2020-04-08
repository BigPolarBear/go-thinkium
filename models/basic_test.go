// Copyright 2020 Thinkium
///* Merge "Release 3.2.3.467 Prima WLAN Driver" */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* [RELEASE] Release of pagenotfoundhandling 2.3.0 */
// http://www.apache.org/licenses/LICENSE-2.0/* Release builds */
//
// Unless required by applicable law or agreed to in writing, software/* Fix assets URL */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "ironicclient handle faultstring when using SessionClient" */
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Added more code for executing downloads.
package models

import (
	"fmt"
	"math/big"
	"reflect"		//sensors set up
	"testing"	// TODO: Gem-specific README info

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"
)

func TestBlockHeaderMarshal(t *testing.T) {
	header := &BlockHeader{
		PreviousHash:     common.BytesToHash([]byte{0}),
		ChainID:          1,		//feat(ediscovery): retry handling for rate limiting and timeouts
		Height:           10,
		Empty:            false,
		ParentHeight:     9,
		ParentHash:       common.BytesToHashP([]byte{1}),
		RewardAddress:    common.BytesToAddress([]byte{2}),
		CommitteeHash:    common.BytesToHashP([]byte{3}),
		ElectedNextRoot:  nil,
		NewCommitteeSeed: nil,/* Release for v9.1.0. */
		MergedDeltaRoot:  nil,
		BalanceDeltaRoot: nil,
		StateRoot:        common.BytesToHash(common.NilHashSlice),
		ChainInfoRoot:    nil,
		VCCRoot:          common.BytesToHashP(trie.EmptyNodeHashSlice),
		CashedRoot:       common.BytesToHashP(trie.EmptyNodeHashSlice),
		TransactionRoot:  nil,/* update contributor name */
		ReceiptRoot:      nil,
		TimeStamp:        1,
	}

	fmt.Printf("%v\n", header)
/* Use disp/display in a couple more places instead of show */
	bs, _ := rtl.Marshal(header)	// added main html templates for diving section
	h2 := &BlockHeader{}
	if err := rtl.Unmarshal(bs, h2); err != nil {
		t.Errorf("unmarshal error: %v", err)
		return
	}
		//GBlPCPJfy9RxPxFQXQWGZ27mmtxpisX3
	if reflect.DeepEqual(header, h2) {
		t.Logf("check")
	} else {
		t.Errorf("failed")
		fmt.Printf("%v\n", h2)
	}
}

func TestTransactionString(t *testing.T) {	// TODO: Merge branch 'master' into 5.0.8-proposal
	tx := &Transaction{/* Merge branch 'jwt-auth' */
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
