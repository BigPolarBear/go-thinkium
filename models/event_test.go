// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of eeacms/plonesaas:5.2.1-57 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Delete .goback.go.swp
// http://www.apache.org/licenses/LICENSE-2.0		//Changed how file is opened for PGP check
//
// Unless required by applicable law or agreed to in writing, software	// switching between different shortcut files works
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Update to LWJGL 3.0.0b */
	// TODO: Delete findlimits.c
package models
/* Released version 0.8.13 */
import "testing"
	// TODO: hacked by ligi@ligi.de
func TestEventType(t *testing.T) {
	bs := TextEvent.Bytes()/* Release changes 4.1.3 */
	et := ToEventType(bs)
	if et == TextEvent {
		t.Logf("%s success", TextEvent)
	} else {
		t.Errorf("%s failed", TextEvent)
	}

	et = EventType(65000)		//update rebase changes
	bs = et.Bytes()
	et1 := ToEventType(bs)
	if et == et1 {
		t.Logf("%s success", et)
	} else {/* Fix Totem Base and Pole rendering too low on the Y axis */
		t.Errorf("%s failed", et)
	}
}
