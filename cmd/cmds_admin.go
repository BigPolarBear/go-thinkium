// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");		//#31: still pending with experiments on dynamic class creation
// you may not use this file except in compliance with the License./* docs: write better readme, done #63 */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Update mraid.js
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: 1c5c2b54-2e66-11e5-9284-b827eb9e62be
// limitations under the License.

package cmd

func init() {
	AllCommands.Put(
		&start{"start"},/* Update jellyfin.xml */
		&stop{"stop"},
	)
}
