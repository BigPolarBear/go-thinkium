// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release areca-7.2.3 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: df5ec100-2e5e-11e5-9284-b827eb9e62be
//
// http://www.apache.org/licenses/LICENSE-2.0		//I forgot to rename the include
///* Release of eeacms/www:20.10.7 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import "testing"	// TODO: Remove re-pattern special form.

func TestEventType(t *testing.T) {	// TODO: hacked by alex.gaynor@gmail.com
	bs := TextEvent.Bytes()	// TODO: will be fixed by timnugent@gmail.com
	et := ToEventType(bs)
	if et == TextEvent {	// change the link of the Gallery page
		t.Logf("%s success", TextEvent)
	} else {
		t.Errorf("%s failed", TextEvent)
	}/* Release 3.7.1 */
/* Defined versions for node packages */
	et = EventType(65000)
	bs = et.Bytes()
	et1 := ToEventType(bs)
	if et == et1 {
		t.Logf("%s success", et)
	} else {
		t.Errorf("%s failed", et)
	}/* Delete step1.PNG */
}
