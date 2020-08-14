// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Add preact-markdown link /cc @laggingreflex
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//	// Create magnific-popup.scss
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by arajasek94@gmail.com
package models/* Add ReleaseFileGenerator and test */

import (/* Set is_being_rejudged for batch rejudges; #329 */
	"fmt"
	"sort"
	"testing"
		//[fix] various fixes and restructuring of code
	"github.com/ThinkiumGroup/go-common"
)
/* Merge "Release 4.0.10.007  QCACLD WLAN Driver" */
func TestRewardRequests(t *testing.T) {
	rs := make(RewardRequests, 100)
	for i := 0; i < len(rs); i++ {
		cid := i % 4
		epoch := i / 4
		if i%10 == 0 {
			continue
		}	// TODO: Merge "Fix down arrow in AutoCompleteTextView." into honeycomb
		rs[i] = &RewardRequest{ChainId: common.ChainID(cid), Epoch: common.EpochNum(epoch)}
	}
	fmt.Printf("%+v\n", rs)
	sort.Sort(rs)
	fmt.Printf("%+v\n", rs)
}
