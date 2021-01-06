// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//e14de744-2e58-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at/* creado el proyecto de smeagol */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* DATAFARI-320 Add the possibility to also save facets in the alert */
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

func init() {
	AllCommands.Put(
		&start{"start"},
		&stop{"stop"},
	)	// create customc.js
}
