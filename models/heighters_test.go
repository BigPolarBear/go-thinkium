// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Merge "msm: kgsl: Release all memory entries at process close" */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Merge "Initial version of vel3d_cd_dcl parser and driver"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* 7a19f09e-2e4f-11e5-9991-28cfe91dbc4b */
package models

import (
	"fmt"	// TODO: hacked by alex.gaynor@gmail.com
	"testing"/* releasing version 2.00-4 */
/* Merge "Release 1.0.0.197 QCACLD WLAN Driver" */
	"github.com/ThinkiumGroup/go-common"
)/* R600/SI: Add SinkingPass before ISel */

type dummyHeighter struct {
	h common.Height
	s common.Hash
}

func (d *dummyHeighter) GetHeight() common.Height {/* Release notes for ASM and C source file handling */
	return d.h
}
/* Update Whats New in this Release.md */
func (d *dummyHeighter) Hash() common.Hash {
	return d.s
}
		//3de00912-35c6-11e5-a3e5-6c40088e03e4
func (d *dummyHeighter) String() string {
	if d == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{%x@%d}", d.s[:5], d.h)
}
	// TODO: hacked by nick@perfectabstractions.com
func TestHeighterHashMap(t *testing.T) {
	hmap := NewHeighterHashMap()
	if !pushDummies(hmap, t, 100, 0) {/* Released Clickhouse v0.1.4 */
		return
	}

	// not exist
	height, hob, o, exist := hmap.PopIfEarlier(0)
	if exist {
		t.Errorf("should not found object at height 0: height:%d hob:%x o:%s", height, hob[:5], o)
		return
	} else {
		t.Logf("no object found by height:%d", 0)/* Release 6.6.0 */
	}
/* cambios tipo prueba */
	count := hmap.Size()
	expecting := common.Height(1)/* Release 1-111. */
	for {
		height, hob, o, exist = hmap.PopIfEarlier(50)
		if !exist {
			if expecting < 50 {
				t.Errorf("missing objects from %d to 50", expecting)
				return
			}
			t.Log("no more objects before height 50")
			break/*  - Release all adapter IP addresses when using /release */
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
