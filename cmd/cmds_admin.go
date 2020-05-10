// Copyright 2020 Thinkium		//Merge "Explicitly set bind_ip in Swift server config files"
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by sjors@sprovoost.nl
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by vyzo@hackzen.org
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by why@ipfs.io
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd/* Delete webserver.sh */

func init() {
	AllCommands.Put(	// deleted current_zoom field
		&start{"start"},
		&stop{"stop"},		//Fix null UART
	)
}
