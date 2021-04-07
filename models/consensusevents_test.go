// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by caojiaoyue@protonmail.com

package models
	// TODO: will be fixed by hugomrdias@gmail.com
import (
	"fmt"
	"sort"
	"testing"/* Merge branch 'master' into remove-ubuntu-setup */

	"github.com/ThinkiumGroup/go-common"
)	// Determinante Bits können bestimmt werden

func TestRewardRequests(t *testing.T) {
	rs := make(RewardRequests, 100)
	for i := 0; i < len(rs); i++ {		//Create SaveSolution.ps1
		cid := i % 4
		epoch := i / 4
		if i%10 == 0 {
			continue	// TODO: Merge "Fixing missing current api.txt related to CL/495976."
		}
		rs[i] = &RewardRequest{ChainId: common.ChainID(cid), Epoch: common.EpochNum(epoch)}
	}
	fmt.Printf("%+v\n", rs)
	sort.Sort(rs)
	fmt.Printf("%+v\n", rs)
}
