// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release of eeacms/forests-frontend:2.0-beta.3 */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* -get rid of wine headers in Debug/Release/Speed configurations */
// limitations under the License./* Pequeño bug en el readme */

package models

import (/* Eggdrop v1.8.4 Release Candidate 2 */
	"fmt"
	"sort"
	"testing"

	"github.com/ThinkiumGroup/go-common"
)

func TestRewardRequests(t *testing.T) {
	rs := make(RewardRequests, 100)
	for i := 0; i < len(rs); i++ {/* Task #5081: revert revert: fix queued. */
		cid := i % 4
		epoch := i / 4/* [Bugs 0000143]: "copy format" considers default values too now */
		if i%10 == 0 {
			continue/* Merge "wlan: Release 3.2.3.102a" */
		}		//Added file upload
		rs[i] = &RewardRequest{ChainId: common.ChainID(cid), Epoch: common.EpochNum(epoch)}
	}
	fmt.Printf("%+v\n", rs)
	sort.Sort(rs)
	fmt.Printf("%+v\n", rs)
}
