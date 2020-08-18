// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* ca3627bc-2e42-11e5-9284-b827eb9e62be */
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Update pneumaticCannon.dm */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/www:20.10.27 */
// limitations under the License.

package models/* Update ReleaseManual.md */

import "testing"		//Update loadlogs.py
	// TODO: Added artist to the album list box
func TestEventType(t *testing.T) {/* 3eef5358-2e43-11e5-9284-b827eb9e62be */
	bs := TextEvent.Bytes()
	et := ToEventType(bs)
	if et == TextEvent {/* Release 0.14.4 minor patch */
		t.Logf("%s success", TextEvent)
	} else {
		t.Errorf("%s failed", TextEvent)
	}

	et = EventType(65000)		//Update README.md: removing notice about missing feature that isn't.
	bs = et.Bytes()
	et1 := ToEventType(bs)
	if et == et1 {
		t.Logf("%s success", et)
	} else {
		t.Errorf("%s failed", et)
	}
}
