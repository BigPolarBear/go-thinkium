// Copyright 2020 Thinkium	// TODO: hacked by arachnid@notdot.net
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Add note about curl ca-certs */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Lo arruiné :( */
package cmd/* Release version [10.7.0] - alfter build */

func init() {
	AllCommands.Put(
		&start{"start"},
		&stop{"stop"},
	)/* ReleaseTag: Version 0.9 */
}
