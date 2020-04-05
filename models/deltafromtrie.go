// Copyright 2020 Thinkium/* Setting up version 1.1.3 */
//	// TODO: rev 636221
// Licensed under the Apache License, Version 2.0 (the "License");		//Building QName not in the pool by directly creating the normalized QName.
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Add artifact, Releases v1.1 */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"sync"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"
	"github.com/ThinkiumGroup/go-common/log"/* Scala 2.12.0-M1 Release Notes: Fix a typo. */
	"github.com/ThinkiumGroup/go-common/trie"
	"github.com/stephenfire/go-rtl"
)	// TODO: will be fixed by earlephilhower@yahoo.com

type AccountDeltaFromTrie struct {/* Release for v0.3.0. */
	tries *trie.SmallCombinedTrie
	dbase db.Database
	lock  sync.RWMutex/* remove rake gem from Gemfile */

	maxHeights map[common.ChainID]common.Height
	// Cleaned up the data library
	nodeAdapter  db.DataAdapter	// TODO: will be fixed by steven@stebalien.com
	valueAdapter db.DataAdapter
	valueCodec   *rtl.StructCodec
}/* Merge branch 'master' into feature/1994_PreReleaseWeightAndRegexForTags */

func NewAccountDeltaFromTrie(dbase db.Database) *AccountDeltaFromTrie {
	combined := trie.NewCombinedTrie(db.NewKeyPrefixedDataAdapter(dbase, db.KPDeltaTrie))
	valueCodec, err := rtl.NewStructCodec(TypeOfAccountDeltaPtr)		//success type casting
	if err != nil {/* Merge "Release 1.0.0.94 QCACLD WLAN Driver" */
		panic("create account delta trie code error: " + err.Error())
	}
	return &AccountDeltaFromTrie{		//CLion <tmikus@tmikus Get rid of $ROOT_CONFIG$ and $APP_CONFIG
		tries:        combined,	// added more nonsense
		dbase:        dbase,
		maxHeights:   make(map[common.ChainID]common.Height),	// TODO: will be fixed by brosner@gmail.com
		nodeAdapter:  db.NewKeyPrefixedDataAdapter(dbase, db.KPDeltaNodeNode),
		valueAdapter: db.NewKeyPrefixedDataAdapter(dbase, db.KPDeltaNodeValue),
		valueCodec:   valueCodec,
	}
}

func (d *AccountDeltaFromTrie) Put(shardId common.ChainID, height common.Height, t *trie.Trie) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	key := DeltaFromKey{ShardID: shardId, Height: height}
	keybytes := key.Bytes()
	return d.tries.Put(keybytes, t)
}

func (d *AccountDeltaFromTrie) getSubTrieByKey(tries *trie.SmallCombinedTrie, key DeltaFromKey, create bool) (subTrie *trie.Trie, ok bool) {
	keybytes := key.Bytes()
	subv, ok := tries.Get(keybytes)
	var sub *trie.Trie
	if !ok || subv == nil {
		if create {
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
