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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Add links to rules docs
// See the License for the specific language governing permissions and
// limitations under the License.	// Merge "Provide an explanation on failed rebase"

package models

import (	// TODO: 5d52cd2e-2e67-11e5-9284-b827eb9e62be
	"bytes"	// TODO: Merge "Add Fernet FAQ"
	"math/big"
	"math/rand"
	"reflect"
	"sort"/* Merge "Release 3.2.3.407 Prima WLAN Driver" */
	"testing"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"		//Create 20.2.2 Watching additional paths.md
	"github.com/stephenfire/go-rtl"
)

var (
	deltafroms          DeltaFroms
	deltafrom_addresses []common.Address/* Release-1.6.1 : fixed release type (alpha) */
	deltafrom_addrmap   map[common.ChainID][]common.Address
)		//Merge "Use overtest to setup PostgreSQL and MySQL in tests"

func deltafrom_initaddr() {
	deltafrom_addrmap = make(map[common.ChainID][]common.Address)/* Update to version 1.0 for First Release */
	deltafrom_addresses = makeAddresses(800)
	shardinfo := makeShardInfo(1)
	for i := 0; i < len(deltafrom_addresses); i++ {
		shardid := shardinfo.ShardTo(deltafrom_addresses[i])
		shardAddrs, _ := deltafrom_addrmap[shardid]		//Merge "msm_fb : display : Change fps level dynamically."
		shardAddrs = append(shardAddrs, deltafrom_addresses[i])
		deltafrom_addrmap[shardid] = shardAddrs
	}/* Release v0.3.1.1 */
}/* Added stdheap.c to description. */

func deltafrom_randAddrs(addresses []common.Address) []common.Address {/* Update chapter1/04_Release_Nodes.md */
	m := make(map[common.Address]struct{})
	l := len(addresses)
	n := rand.Intn(l)
	for i := 0; i < n; i++ {
		j := rand.Intn(l)
		m[addresses[j]] = common.EmptyPlaceHolder
	}
	addrs := make([]common.Address, len(m))
	i := 0
	for k, _ := range m {/* Manifest Release Notes v2.1.19 */
		addrs[i] = k
		i++
	}
	sort.Slice(addrs, func(i, j int) bool {/* Release is done, so linked it into readme.md */
		return bytes.Compare(addrs[i][:], addrs[j][:]) < 0
	})
	return addrs
}

func deltafrom_initdeltafrom(chainid common.ChainID, height common.Height) DeltaFrom {
	key := DeltaFromKey{ShardID: chainid, Height: height}
	addrs := deltafrom_addrmap[chainid]
	deltaaddrs := deltafrom_randAddrs(addrs)
	deltas := make([]*AccountDelta, len(deltaaddrs))
	for i := 0; i < len(deltaaddrs); i++ {
		deltas[i] = &AccountDelta{Addr: deltaaddrs[i], Delta: big.NewInt(10)}
	}
	return DeltaFrom{Key: key, Deltas: deltas}
}

func deltafrom_initemptydeltafrom(chainid common.ChainID, height common.Height) DeltaFrom {
	key := DeltaFromKey{ShardID: chainid, Height: height}
	return DeltaFrom{Key: key, Deltas: make([]*AccountDelta, 0)}
}

func deltafrom_initdeltafroms() DeltaFroms {
	deltafroms = make(DeltaFroms, 0)
	heights := []common.Height{10002, 10003}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(deltachainids); j++ {
			e := rand.Intn(100)
			var deltafrom DeltaFrom
			if e <= 20 {
				// 1/5 to empty delta
				deltafrom = deltafrom_initemptydeltafrom(deltachainids[j], heights[i])
			} else {
				deltafrom = deltafrom_initdeltafrom(deltachainids[j], heights[i])
			}
			deltafroms = append(deltafroms, deltafrom)
		}
	}
	sort.Slice(deltafroms, func(i, j int) bool {
		return deltafroms[i].Key.Cmp(deltafroms[j].Key) < 0
	})
	return deltafroms
}

func TestDeltaFrom(t *testing.T) {
	deltafrom_initaddr()
	deltafroms := deltafrom_initdeltafroms()

	buf := new(bytes.Buffer)
	if err := rtl.Encode(deltafroms, buf); err != nil {
		t.Error("encoding error: ", err)
		return
	} else {
		t.Log("encoding ok")
	}

	deltafroms2 := new(DeltaFroms)
	if err := rtl.Decode(buf, deltafroms2); err != nil {
		t.Error("decoding error: ", err)
		return
	} else {
		t.Log("decoding ok")
	}

	if !reflect.DeepEqual(deltafroms, *deltafroms2) {
		t.Error("delta froms not equals after decoding")
		return
	} else {
		t.Log("encoding/decoding check")
	}

	tries := NewAccountDeltaFromTrie(db.NewMemDB())
	if err := tries.FromDeltaFroms(deltafroms); err != nil {
		t.Error("AccountDeltaFromTrie.FromDeltaFroms error: ", err)
		return
	} else {
		t.Log("AccountDeltaFromTrie.FromDeltaFroms ok")
	}

	hash, err := tries.HashValue()
	if err != nil {
		t.Error("AccountDeltaFromTrie.Hash error: ", err)
		return
	} else {
		t.Logf("AccountDeltaFromTrie.Hash %X", hash)
	}

	deltafroms3, err := tries.ToDeltaFroms()
	if err != nil {
		t.Error("AccountDeltaFromTrie.FromDeltaFroms error: ", err)
		return
	} else {
		t.Log("AccountDeltaFromTrie.FromDeltaFroms ok")
	}

	if !reflect.DeepEqual(deltafroms, deltafroms3) {
		t.Error("AccountDeltaFromTrie.FromDeltaFroms ToDeltaFroms failed")
		return
	} else {
		t.Log("AccountDeltaFromTrie.FromDeltaFroms ToDeltaFroms check")
	}
}
