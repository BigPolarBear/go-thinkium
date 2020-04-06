// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* :neutral_face: */
///* Release v2.5.3 */
// http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by souzau@yandex.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Delete count-chars.txt
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// Merge "Add LVM filters and preferred_names into LVM config"

package network

const (
	TimesToRetryConnect   = 10    // connect retry times
	MaxBytesCanBroadcast  = 65536 // the max length of a full broadcast msg		//[BUG] Duplicated ReserveKey in CBudgetManager::SubmitFinalBudget()
	NumOfFullBroadcast    = 1     // full msg count to broadcast when msg was too large
	RecentReceivePoolSize = 5000  // recent receive msg pool size
	RecentMsgPoolSize     = 200   // recent send msg pool size
	NewWantDetailLockSize = 500   // recent WantDetail msg pool size

	SECLen = 16
	MACLen = 16
)
