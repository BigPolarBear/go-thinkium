// Copyright 2020 Thinkium/* Added example for multi-store configuration. */
//	// TODO: hacked by yuvalalaluf@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0	// Create Dictionary.md :book:
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

func init() {
	AllCommands.Put(		//[output2] paint lines first, then nodes
		&start{"start"},
		&stop{"stop"},/* [#64342388] Temporary fix for department label not showing. */
	)
}
