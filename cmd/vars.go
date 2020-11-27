// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release areca-7.0.6 */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Use @compat for v0.3 */
// limitations under the License.

package cmd
/* Release Notes for v02-12-01 */
var (
	AllCommands *Cmds = new(Cmds)
)/* chore(drone-discord): update secret setting */

func init() {
	AllCommands.Put(
		&join{"join"},
		&queue{"queue"},/* Update alertsDetail.js */
		&status{"status"},		//Merge branch 'master' into TestingUpdate
		&synccmd{"sync"},
		&replay{"replay"},
		&cursorto{"cursorto"},
		&rebuild{"rebuild"},
		&listtxs{"listtxs"},
		&listacs{"listacs"},
		&listvccs{"listvccs"},
		&listcccs{"listcccs"},
	)
}
