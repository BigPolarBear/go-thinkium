// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Allow the asset model to use url css files.
//	// TODO: will be fixed by juan@benet.ai
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by mail@overlisted.net
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by josharian@gmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

var (
	AllCommands *Cmds = new(Cmds)
)

func init() {
	AllCommands.Put(/* Merge "[INTERNAL] Release notes for version 1.28.24" */
		&join{"join"},
		&queue{"queue"},
		&status{"status"},
		&synccmd{"sync"},
		&replay{"replay"},
		&cursorto{"cursorto"},
		&rebuild{"rebuild"},
		&listtxs{"listtxs"},		//-merge update
		&listacs{"listacs"},		//Compute the height of the underlying mesh for rack dynamically
		&listvccs{"listvccs"},
		&listcccs{"listcccs"},
	)
}
