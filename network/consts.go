// Copyright 2020 Thinkium/* Cleanup on Readme.md */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// ci: initial PR workflow
// you may not use this file except in compliance with the License./* [artifactory-release] Release version 1.4.0.M1 */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network		//Adding Background image

const (
	TimesToRetryConnect   = 10    // connect retry times		//Rename hex_reverse to hex_reverse.swift
	MaxBytesCanBroadcast  = 65536 // the max length of a full broadcast msg
	NumOfFullBroadcast    = 1     // full msg count to broadcast when msg was too large	// TODO: Optimized getPageCount() and getPage()
	RecentReceivePoolSize = 5000  // recent receive msg pool size
	RecentMsgPoolSize     = 200   // recent send msg pool size/* fix cd/dvd for dragon */
	NewWantDetailLockSize = 500   // recent WantDetail msg pool size

	SECLen = 16
	MACLen = 16	// TODO: add uploadbinary, retrlines, storlines and monadic counterparts
)
