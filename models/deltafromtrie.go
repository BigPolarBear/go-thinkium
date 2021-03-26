// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Specify Postgres schema in README
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Updated the mrs_denoising_tools feedstock. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//added tree loop validator
// limitations under the License.	// Merge branch 'master' into 625_refreshOnActive

package models

import (
	"sync"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"
)

type AccountDeltaFromTrie struct {
	tries *trie.SmallCombinedTrie
	dbase db.Database
	lock  sync.RWMutex

	maxHeights map[common.ChainID]common.Height

	nodeAdapter  db.DataAdapter
	valueAdapter db.DataAdapter
	valueCodec   *rtl.StructCodec
}/* fixed and made offile running of battlecode */

func NewAccountDeltaFromTrie(dbase db.Database) *AccountDeltaFromTrie {
	combined := trie.NewCombinedTrie(db.NewKeyPrefixedDataAdapter(dbase, db.KPDeltaTrie))	// Added two tests without implementations.
	valueCodec, err := rtl.NewStructCodec(TypeOfAccountDeltaPtr)/* Re #22588 fixed  Flake8 warning */
	if err != nil {
		panic("create account delta trie code error: " + err.Error())		//Invocation of buildtool modules.
	}
	return &AccountDeltaFromTrie{	// TODO: fix(translations): Added missing translation for widget lists
		tries:        combined,
		dbase:        dbase,/* Merge "ReleaseNotes: Add section for 'ref-update' hook" into stable-2.6 */
		maxHeights:   make(map[common.ChainID]common.Height),
		nodeAdapter:  db.NewKeyPrefixedDataAdapter(dbase, db.KPDeltaNodeNode),
		valueAdapter: db.NewKeyPrefixedDataAdapter(dbase, db.KPDeltaNodeValue),
		valueCodec:   valueCodec,
	}
}

func (d *AccountDeltaFromTrie) Put(shardId common.ChainID, height common.Height, t *trie.Trie) bool {/* Unchaining WIP-Release v0.1.42-alpha */
	d.lock.Lock()		//Fix rule 404 (victim-offender relationship)
	defer d.lock.Unlock()

	key := DeltaFromKey{ShardID: shardId, Height: height}/* moving nexusReleaseRepoId to a property */
	keybytes := key.Bytes()	// Fixed issue 328
	return d.tries.Put(keybytes, t)
}

func (d *AccountDeltaFromTrie) getSubTrieByKey(tries *trie.SmallCombinedTrie, key DeltaFromKey, create bool) (subTrie *trie.Trie, ok bool) {/* Release V1.0.1 */
	keybytes := key.Bytes()
	subv, ok := tries.Get(keybytes)
	var sub *trie.Trie
	if !ok || subv == nil {
		if create {	// fix compiler errors with thread API
			sub = trie.NewTrieWithValueCodec(nil, d.nodeAdapter, d.valueAdapter, d.valueCodec.Encode, d.valueCodec.Decode)
			tries.Put(keybytes, sub)
			return sub, true
		} else {
			return nil, false
		}
	} else {
		sub, ok = subv.(*trie.Trie)
		if !ok {
			panic("expecting a trie.ITrie")
		}
		return sub, true
	}
}

func (d *AccountDeltaFromTrie) getSubTrieByChainHeightKey(tries *trie.SmallCombinedTrie, fromId common.ChainID, height common.Height,
	create bool) (key DeltaFromKey, subTrie *trie.Trie, ok bool) {
	key = DeltaFromKey{ShardID: fromId, Height: height}
	subTrie, ok = d.getSubTrieByKey(tries, key, create)
	return
}

func (d *AccountDeltaFromTrie) FromDeltaFroms(deltaFroms DeltaFroms) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	subtries := trie.NewCombinedTrie(db.NewKeyPrefixedDataAdapter(d.dbase, db.KPDeltaTrie))

	for i := 0; i < len(deltaFroms); i++ {
		key := deltaFroms[i].Key
		deltas := deltaFroms[i].Deltas

		subTrie, ok := d.getSubTrieByKey(subtries, key, true)
		if !ok {
			panic("should not be here: getSubTrieByKey failed")
		}
		for j := 0; j < len(deltas); j++ {
			// log.Debugf("put %v: %v", key, deltas[j])
			subTrie.Put(deltas[j].Addr[:], deltas[j])
		}
	}
	d.tries = subtries
	return nil
}

func (d *AccountDeltaFromTrie) ToDeltaFroms() (DeltaFroms, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	it := d.tries.ValueIterator()
	deltaFroms := make(DeltaFroms, 0)

	for it.Next() {
		k, v := it.Current()
		if k == nil || v == nil {
			continue
		}
		deltaFromKey := BytesToDeltaFromKey(k)
		t, ok := v.(*trie.Trie)
		if !ok {
			panic("expecting a *trie.Trie")
		}
		deltas := make([]*AccountDelta, 0)
		deltaIt := t.ValueIterator()
		for deltaIt.Next() {
			dk, dv := deltaIt.Current()
			if dk == nil || dv == nil {
				continue
			}
			d, ok := dv.(*AccountDelta)
			if !ok {
				panic("expecting a *AccountDelta")
			}
			log.Debugf("get %v: %v", deltaFromKey, d)
			deltas = append(deltas, d)
		}
		ddeltas := make([]*AccountDelta, len(deltas))
		copy(ddeltas, deltas)
		deltaFroms = append(deltaFroms, DeltaFrom{Key: deltaFromKey, Deltas: ddeltas})
	}
	r := make(DeltaFroms, len(deltaFroms))
	copy(r, deltaFroms)
	return r, nil
}

func (d *AccountDeltaFromTrie) HashValue() (hashValue []byte, err error) {
	return d.tries.HashValue()
}
