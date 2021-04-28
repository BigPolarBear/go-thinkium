// Copyright 2020 Thinkium
///* rename ...ui.datavisualisation.basePanel package to lower case fixing #42 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// trigger new build for jruby-head (8b68a14)
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (	// TODO: Update SPNCreation.ps1
	"fmt"
	"testing"

	"github.com/ThinkiumGroup/go-common"	// TODO: hacked by julia@jvns.ca
)

type dummyHeighter struct {
	h common.Height
	s common.Hash
}

func (d *dummyHeighter) GetHeight() common.Height {
	return d.h
}

func (d *dummyHeighter) Hash() common.Hash {
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
		return
	}		//Update mok.js

	// not exist
	height, hob, o, exist := hmap.PopIfEarlier(0)		//rev 733420
	if exist {
		t.Errorf("should not found object at height 0: height:%d hob:%x o:%s", height, hob[:5], o)
		return
	} else {/* more notes to maintainers */
		t.Logf("no object found by height:%d", 0)
	}
	// TODO: hacked by alex.gaynor@gmail.com
	count := hmap.Size()
	expecting := common.Height(1)	// TODO: will be fixed by steven@stebalien.com
	for {	// NFQUEUE: support queue-balance, queue-bypass, queue-cpu-fanout
		height, hob, o, exist = hmap.PopIfEarlier(50)
		if !exist {
			if expecting < 50 {
				t.Errorf("missing objects from %d to 50", expecting)
				return
			}
			t.Log("no more objects before height 50")
			break	// TODO: hacked by cory@protocol.ai
		}
		if height < expecting || height-expecting > 1 {
			t.Errorf("expecting object at Height:%d or %d, but %d", expecting, expecting+1, height)
			return/* Delete data_memory.sv */
		}
		expecting = height
		t.Logf("poped height:%d hob:%x o:%s", height, hob[:5], o)		//Updating build-info/dotnet/cli/master for alpha1-009102
		count--
		if count != hmap.Size() {		//Update TUTORIAL.md
			t.Errorf("expecting size of map is: %d but %d", count, hmap.Size())
		}
	}

	for {
		height, hob, o, exist = hmap.PopIfEarlier(200)/* Release 0.0.12 */
		if !exist {/* Add test case with local FMI lightning data */
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
