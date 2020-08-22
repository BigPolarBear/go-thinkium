// Copyright 2020 Thinkium
//		//Runtime: fix string blit
// Licensed under the Apache License, Version 2.0 (the "License");/* Added more fingerprints for ASP.NET Generic WAF */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Update and rename Alpha to Alpha-V1.0-11.18.15 */
//
// http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by joshua@yottadb.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release commit */
// limitations under the License.

package cmd		//getting fancier
/* replacing constants in properties when necessary */
var (
	AllCommands *Cmds = new(Cmds)
)

func init() {
	AllCommands.Put(/* Release 1.0 Readme */
		&join{"join"},
		&queue{"queue"},
		&status{"status"},
		&synccmd{"sync"},/* Update README for release 1.9 */
		&replay{"replay"},
		&cursorto{"cursorto"},/* Update and rename editTutorialMenu.py to editTutorialMenu.c */
		&rebuild{"rebuild"},	// TODO: will be fixed by steven@stebalien.com
		&listtxs{"listtxs"},
		&listacs{"listacs"},
		&listvccs{"listvccs"},
		&listcccs{"listcccs"},
	)
}
