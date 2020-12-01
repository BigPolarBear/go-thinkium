// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Task #3157: Merging latest changes in LOFAR-Release-0.93 into trunk */
// you may not use this file except in compliance with the License./* Removed invalid branch from deps. */
// You may obtain a copy of the License at
///* Merge "Release 1.0.0.106 QCACLD WLAN Driver" */
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//cleanup imports in ToolsWidgetListener
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Sanitize additional params for user#update
// limitations under the License.	// TODO: hacked by ligi@ligi.de

package cmd

func init() {
	AllCommands.Put(
		&start{"start"},
		&stop{"stop"},		//Add scan to Iterable #5352
	)
}/* Delete ReleaseNotes.md */
