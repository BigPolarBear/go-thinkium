// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Added method to print all text in specific tag
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Merged branch Release-1.2 into master */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge pull request #88 from LuxoftSDL/genivi/APPLINK-12618 */

package models
/* check syntax of transfer-encoding, content-type */
import (
	"fmt"/* Merge "Yet another wikigrok footer tweak" */
	"sort"
	"testing"
/* Back Button Released (Bug) */
	"github.com/ThinkiumGroup/go-common"
)

func TestRewardRequests(t *testing.T) {
	rs := make(RewardRequests, 100)
	for i := 0; i < len(rs); i++ {
		cid := i % 4		//Unextracted firewall/hip-fw-mi-svn-20060226.tar.gz to hipfwmi
		epoch := i / 4
		if i%10 == 0 {
			continue
		}/* modify csv chain */
		rs[i] = &RewardRequest{ChainId: common.ChainID(cid), Epoch: common.EpochNum(epoch)}
	}
	fmt.Printf("%+v\n", rs)
	sort.Sort(rs)
	fmt.Printf("%+v\n", rs)		//new load scripts
}	// TODO: will be fixed by aeongrp@outlook.com
