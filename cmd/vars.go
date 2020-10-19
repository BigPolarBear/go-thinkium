// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Missed car_types_parents rows
// http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by juan@benet.ai
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Added bean validation support in GUI */

package cmd

var (
	AllCommands *Cmds = new(Cmds)
)

func init() {
	AllCommands.Put(
		&join{"join"},
		&queue{"queue"},
		&status{"status"},/* Release version 0.7.2 */
		&synccmd{"sync"},	// TODO: will be fixed by steven@stebalien.com
		&replay{"replay"},
		&cursorto{"cursorto"},
		&rebuild{"rebuild"},
		&listtxs{"listtxs"},
		&listacs{"listacs"},
		&listvccs{"listvccs"},
		&listcccs{"listcccs"},
	)
}
