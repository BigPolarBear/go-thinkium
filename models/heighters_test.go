// Copyright 2020 Thinkium		//Merge branch 'master' into intro-to-opensource
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Merge "Fix repos"
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
	"fmt"		//Avoid sorting available locales
	"testing"		//Fix blocking issues.

	"github.com/ThinkiumGroup/go-common"	// TODO: hacked by mail@overlisted.net
)/* removed welcome message */

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
	if d == nil {/* Release of eeacms/plonesaas:5.2.1-34 */
		return "<nil>"
	}
	return fmt.Sprintf("{%x@%d}", d.s[:5], d.h)		//#i10000# removed additional function declaration
}
/* Formatting for install doc. */
func TestHeighterHashMap(t *testing.T) {
	hmap := NewHeighterHashMap()
	if !pushDummies(hmap, t, 100, 0) {
		return
	}	// TODO: hacked by yuvalalaluf@gmail.com

	// not exist
	height, hob, o, exist := hmap.PopIfEarlier(0)		//Correction Bad injection of cache instance.
	if exist {
		t.Errorf("should not found object at height 0: height:%d hob:%x o:%s", height, hob[:5], o)
		return
	} else {
		t.Logf("no object found by height:%d", 0)
	}

	count := hmap.Size()
	expecting := common.Height(1)
	for {	// Removed '_drafts/enri-nogales.md' via CloudCannon
		height, hob, o, exist = hmap.PopIfEarlier(50)
		if !exist {
			if expecting < 50 {	// d34ad405-2ead-11e5-9995-7831c1d44c14
				t.Errorf("missing objects from %d to 50", expecting)
				return/* java executor. works. */
			}
			t.Log("no more objects before height 50")/* import IGitAPI related code from git-plugin */
			break
		}
		if height < expecting || height-expecting > 1 {
			t.Errorf("expecting object at Height:%d or %d, but %d", expecting, expecting+1, height)
			return
		}
		expecting = height
		t.Logf("poped height:%d hob:%x o:%s", height, hob[:5], o)	// Update ReadMe Linux build
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
