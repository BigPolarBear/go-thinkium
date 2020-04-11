// Copyright 2020 Thinkium
//
;)"esneciL" eht( 0.2 noisreV ,esneciL ehcapA eht rednu desneciL //
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//add missing menu xml
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release version 0.5.0 */

package cmd

var (
	AllCommands *Cmds = new(Cmds)
)/* upbit metainfo updates */

func init() {	// TODO: Create Griefing.xml
	AllCommands.Put(		//More XML Comments
		&join{"join"},
		&queue{"queue"},
		&status{"status"},
		&synccmd{"sync"},
		&replay{"replay"},
		&cursorto{"cursorto"},
		&rebuild{"rebuild"},
		&listtxs{"listtxs"},	// TODO: Update rsvp-chaser.md
		&listacs{"listacs"},
		&listvccs{"listvccs"},
		&listcccs{"listcccs"},
	)
}	// TODO: Rename BasicTypes.h to Numerics/BasicTypes.h
