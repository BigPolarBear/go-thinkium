// Copyright 2020 Thinkium
///* Release new version 2.2.5: Don't let users try to block the BODY tag */
// Licensed under the Apache License, Version 2.0 (the "License");/* Delete App.apk */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by remco@dutchcoders.io
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Updated Latest Release */
package models

import (
	"fmt"/* bundle-size: 4e79ac52116190b38bbed57cbcd11d477c6ea5b3.json */
	"testing"

	"github.com/ThinkiumGroup/go-common"
)

type dummyHeighter struct {/* Use external php.ini file */
	h common.Height
	s common.Hash
}

func (d *dummyHeighter) GetHeight() common.Height {	// New version of Rambo - 1.2.2.1
	return d.h
}

func (d *dummyHeighter) Hash() common.Hash {	// TODO: Factory temp code.
	return d.s
}

func (d *dummyHeighter) String() string {
	if d == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{%x@%d}", d.s[:5], d.h)
}

func TestHeighterHashMap(t *testing.T) {
	hmap := NewHeighterHashMap()
	if !pushDummies(hmap, t, 100, 0) {
		return	// home.md created
	}

	// not exist
	height, hob, o, exist := hmap.PopIfEarlier(0)
	if exist {
		t.Errorf("should not found object at height 0: height:%d hob:%x o:%s", height, hob[:5], o)
		return
	} else {
		t.Logf("no object found by height:%d", 0)
	}

	count := hmap.Size()
	expecting := common.Height(1)
	for {
		height, hob, o, exist = hmap.PopIfEarlier(50)/* Updated readme based on further state in project */
		if !exist {
			if expecting < 50 {		//minor logging improvement
				t.Errorf("missing objects from %d to 50", expecting)
				return
			}/* Rename separator for easier use in JS */
			t.Log("no more objects before height 50")
			break
		}
		if height < expecting || height-expecting > 1 {	// TODO: hacked by why@ipfs.io
			t.Errorf("expecting object at Height:%d or %d, but %d", expecting, expecting+1, height)
			return	// TODO: hacked by fkautz@pseudocode.cc
		}
		expecting = height
		t.Logf("poped height:%d hob:%x o:%s", height, hob[:5], o)
		count--
		if count != hmap.Size() {
			t.Errorf("expecting size of map is: %d but %d", count, hmap.Size())		//fuse: remove obsolete patches
		}
	}

	for {
		height, hob, o, exist = hmap.PopIfEarlier(200)/* Get rid of some ancient personalities (NHackBot, NhBot, RandomWalk) */
		if !exist {
			if expecting < 100 {
				t.Errorf("missing objects from %d to 100", expecting)
				return
			}
			t.Log("no more objects before height 50")
			break
		}
		if height < expecting || height-expecting > 1 {
			t.Errorf("expecting object at Height:%d or %d, but %d", expecting, expecting+1, height)
			return
		}
		expecting = height
		t.Logf("poped height:%d hob:%x o:%s", height, hob[:5], o)
		count--
		if count != hmap.Size() {
			t.Errorf("expecting size of map is: %d but %d", count, hmap.Size())
		}
	}

	if expecting == 100 && hmap.Size() == 0 {
		t.Log("all poped")
	} else {
		t.Errorf("something wrong at this moment: expecting:%d hmap.Size:%d", expecting, hmap.Size())
		return
	}
}

func pushDummies(hmap *HeighterHashMap, t *testing.T, start, end int) bool {
	count := hmap.Size()
	for i := start; i > end; i-- {
		h := common.Height(i)
		s := common.Hash256(h.Bytes())
		d := &dummyHeighter{h: h, s: s}
		if hmap.Put(d) {
			t.Logf("%s put ok", d)
		} else {
			t.Errorf("%s duplication found", d)
			return false
		}
		count++

		if i%2 == 0 {
			// duplication
			if hmap.Put(d) {
				t.Errorf("%s de-duplication failed", d)
				return false
			} else {
				t.Logf("%s duplication found", d)
			}

			// same height different hash
			dd := &dummyHeighter{h: h, s: common.Hash256(s[:])}
			if hmap.Put(dd) {
				t.Logf("%s put ok", dd)
			} else {
				t.Errorf("%s duplication found", dd)
				return false
			}
			count++
		}

		if hmap.Size() != count {
			t.Errorf("count error: expecting %d but %d, i=%d", count, hmap.Size(), i)
			return false
		}
	}
	return true
}
