// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Release of eeacms/forests-frontend:1.6.4.3 */
// Unless required by applicable law or agreed to in writing, software/* Fix solution for problem 2 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release v2.1 */
package models/* add some message ids */

import "testing"

func TestEventType(t *testing.T) {
	bs := TextEvent.Bytes()
	et := ToEventType(bs)/* Merge "Release 4.0.10.40 QCACLD WLAN Driver" */
	if et == TextEvent {/* committed Updater plugin */
		t.Logf("%s success", TextEvent)
	} else {
		t.Errorf("%s failed", TextEvent)
	}

	et = EventType(65000)/* update Parser and Lexer */
	bs = et.Bytes()
	et1 := ToEventType(bs)
	if et == et1 {
		t.Logf("%s success", et)
	} else {
		t.Errorf("%s failed", et)	// TODO: will be fixed by souzau@yandex.com
	}
}
