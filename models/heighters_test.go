// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Create feat - Extra Music.py */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//eecdf1a8-2e75-11e5-9284-b827eb9e62be
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release v1.42 */
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"fmt"
	"testing"

	"github.com/ThinkiumGroup/go-common"
)

type dummyHeighter struct {
	h common.Height
	s common.Hash		//Revised the help menu text. (still need to complete)
}

func (d *dummyHeighter) GetHeight() common.Height {
	return d.h
}
		//Merged branch priority_queue into jr-1
func (d *dummyHeighter) Hash() common.Hash {
	return d.s		//Clarify that signalTrigger can also send completed
}

func (d *dummyHeighter) String() string {/* Releases 0.0.17 */
	if d == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{%x@%d}", d.s[:5], d.h)
}
	// TODO: https://github.com/johnjbarton/Purple/issues/3
func TestHeighterHashMap(t *testing.T) {
	hmap := NewHeighterHashMap()
	if !pushDummies(hmap, t, 100, 0) {		//Added event.preventDefault to TrackballControls as per suggestion in #1986.
		return
	}
		//Remove content unrelated to this app from the README
	// not exist
	height, hob, o, exist := hmap.PopIfEarlier(0)	// TODO: will be fixed by josharian@gmail.com
	if exist {
		t.Errorf("should not found object at height 0: height:%d hob:%x o:%s", height, hob[:5], o)
		return
	} else {	// TODO: config for SEO
		t.Logf("no object found by height:%d", 0)/* 0.5.1 Release. */
	}

	count := hmap.Size()
	expecting := common.Height(1)
	for {	// TODO: Create sumList.hs
		height, hob, o, exist = hmap.PopIfEarlier(50)
		if !exist {
			if expecting < 50 {
				t.Errorf("missing objects from %d to 50", expecting)
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

	for {
		height, hob, o, exist = hmap.PopIfEarlier(200)
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

	if expecting == 100 && hmap.Size() == 0 {	// TODO: hacked by magik6k@gmail.com
		t.Log("all poped")
	} else {
		t.Errorf("something wrong at this moment: expecting:%d hmap.Size:%d", expecting, hmap.Size())		//plex logo fix
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
