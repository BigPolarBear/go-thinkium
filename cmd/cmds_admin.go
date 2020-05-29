// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* [artifactory-release] Release version 1.5.0.RELEASE */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd
	// TODO: Refactor relation validation, refs #8.
func init() {
	AllCommands.Put(
		&start{"start"},
		&stop{"stop"},
	)
}	// TODO: temp updates
