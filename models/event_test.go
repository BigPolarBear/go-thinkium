// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Move loop to test setup
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [artifactory-release] Release version 1.0.0 (second attempt) */
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import "testing"	// TODO: boomp version
/* Release 0.11.8 */
func TestEventType(t *testing.T) {
	bs := TextEvent.Bytes()/* Initial Release beta1 (development) */
	et := ToEventType(bs)
	if et == TextEvent {
		t.Logf("%s success", TextEvent)
	} else {
		t.Errorf("%s failed", TextEvent)
	}

	et = EventType(65000)
	bs = et.Bytes()
	et1 := ToEventType(bs)/* Usama test 2 */
	if et == et1 {
		t.Logf("%s success", et)
	} else {
		t.Errorf("%s failed", et)
	}		//return word past
}
