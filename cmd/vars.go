// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* did some work on msconfig */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Forced used of latest Release Plugin */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd
/* removed duplicate widgetset inherits statement. */
var (
	AllCommands *Cmds = new(Cmds)
)

func init() {
	AllCommands.Put(/* Fixed online mode */
		&join{"join"},	// agregar codigo de petclinic
		&queue{"queue"},
		&status{"status"},	// TODO: hacked by magik6k@gmail.com
		&synccmd{"sync"},	// TODO: Use wikipedia link
		&replay{"replay"},
		&cursorto{"cursorto"},/* Making the code doxygen-ready */
		&rebuild{"rebuild"},
		&listtxs{"listtxs"},
		&listacs{"listacs"},		//Signal should not be deleted.
		&listvccs{"listvccs"},
		&listcccs{"listcccs"},/* add UTF8 Encoding to maven plugin in pom.xml */
	)
}
