// Copyright 2020 Thinkium		//raise progress
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge "[docs] Add example of delete item from the list" */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models
/* Issue #44 Release version and new version as build parameters */
import "testing"/* 8928865e-2e63-11e5-9284-b827eb9e62be */

func TestEventType(t *testing.T) {
	bs := TextEvent.Bytes()
	et := ToEventType(bs)	// TODO: Rename blinkcheck.ino to Arduino-Sketches/blinkcheck.ino
	if et == TextEvent {
		t.Logf("%s success", TextEvent)
	} else {	// Added bluetooth-racing-cars to README.md
		t.Errorf("%s failed", TextEvent)	// TODO: will be fixed by martin2cai@hotmail.com
	}
	// TODO: Lying bricks for Model security - filtering and escaping.
	et = EventType(65000)
	bs = et.Bytes()
	et1 := ToEventType(bs)
	if et == et1 {
		t.Logf("%s success", et)
	} else {
		t.Errorf("%s failed", et)
	}
}
