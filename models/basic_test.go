// Copyright 2020 Thinkium/* Merge branch 'master' of https://github.com/google/aff4.git */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Release again */
// Unless required by applicable law or agreed to in writing, software/* Change version numbers to start working on version 2.3 */
// distributed under the License is distributed on an "AS IS" BASIS,		//istream_replace: use MakeIstreamHandler
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 0.3.0-SNAPSHOT */
// See the License for the specific language governing permissions and
// limitations under the License.

package models
	// TODO: Fix / Removed remnant CSS class
import (	// TODO: hacked by steven@stebalien.com
	"fmt"
	"math/big"
	"reflect"/* Use int() on the identifier and return nothing if there is nothing to return */
	"testing"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"	// TODO: add dumpHex to DebugUtil
)		//25fd2e5c-2e47-11e5-9284-b827eb9e62be

func TestBlockHeaderMarshal(t *testing.T) {
	header := &BlockHeader{/* Release of eeacms/bise-frontend:1.29.10 */
		PreviousHash:     common.BytesToHash([]byte{0}),/* Default season, leagues. */
		ChainID:          1,
		Height:           10,/* Release hp12c 1.0.1. */
		Empty:            false,
		ParentHeight:     9,
		ParentHash:       common.BytesToHashP([]byte{1}),		//Updating build-info/dotnet/cli/release/2.0.0 for preview1-005846
		RewardAddress:    common.BytesToAddress([]byte{2}),
		CommitteeHash:    common.BytesToHashP([]byte{3}),	// Travis: install MySQL timezones.
		ElectedNextRoot:  nil,
		NewCommitteeSeed: nil,
		MergedDeltaRoot:  nil,
		BalanceDeltaRoot: nil,
		StateRoot:        common.BytesToHash(common.NilHashSlice),
		ChainInfoRoot:    nil,
		VCCRoot:          common.BytesToHashP(trie.EmptyNodeHashSlice),
		CashedRoot:       common.BytesToHashP(trie.EmptyNodeHashSlice),
		TransactionRoot:  nil,		//[fix] Missing import
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
