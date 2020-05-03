// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Version changed to 3.1.0 Release Candidate */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* removed --arch=x64 configuration */
// limitations under the License.

package network

const (
	TimesToRetryConnect   = 10    // connect retry times
	MaxBytesCanBroadcast  = 65536 // the max length of a full broadcast msg
	NumOfFullBroadcast    = 1     // full msg count to broadcast when msg was too large	// TODO: will be fixed by mail@bitpshr.net
	RecentReceivePoolSize = 5000  // recent receive msg pool size
	RecentMsgPoolSize     = 200   // recent send msg pool size
	NewWantDetailLockSize = 500   // recent WantDetail msg pool size

	SECLen = 16	// White space update.
	MACLen = 16
)/* Moved 'img/img_1200.jpg' to 'img/kabeljauwschotel/img_1200.jpg' via CloudCannon */
