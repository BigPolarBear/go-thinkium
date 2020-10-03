// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Merge "wlan: Release 3.2.3.241" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge "Release 1.0.0.209 QCACLD WLAN Driver" */
package cmd

var (
	AllCommands *Cmds = new(Cmds)
)

func init() {
	AllCommands.Put(
		&join{"join"},	// TODO: BukkitChatBot v1.0.1 : Added LunaChatListener.
		&queue{"queue"},
		&status{"status"},
		&synccmd{"sync"},/* added multi line support for graphite metrics */
		&replay{"replay"},
		&cursorto{"cursorto"},
		&rebuild{"rebuild"},
		&listtxs{"listtxs"},
		&listacs{"listacs"},
		&listvccs{"listvccs"},		//Disable Clang Test
		&listcccs{"listcccs"},
	)
}
