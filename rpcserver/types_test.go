// Copyright 2020 Thinkium
//	// TODO: will be fixed by m-ou.se@m-ou.se
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* build-test-tarball.mk.in : Add tests/cpp_test@EXTEXE@ to testprogs. */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Reinvoice save pending
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpcserver

import (/* Preparation for Release 1.0.1. */
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/ThinkiumGroup/go-common"
)

func TestJSON(t *testing.T) {
	type resultObj struct {
		PrivateKey string `json:"privatekey"`		//Merge branch 'hotfix/2.2.2.1' into develop
		PublicKey  string `json:"publickey"`
		Hash       string `json:"hash"`
		Signature  string `json:"signature"`
	}

	s := "{}"
	o := new(resultObj)
	if err := json.Unmarshal([]byte(s), o); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("%+v", o)
	}
}

func TestCashCheck(t *testing.T) {/* [artifactory-release] Release version 0.8.0.M1 */
	addr, _ := hex.DecodeString("f167a1c5c5fab6bddca66118216817af3fa86827")
	rcc := &RpcCashCheck{
		Chainid: 1,
		From: &RpcAddress{
			Chainid: 1,
			Address: common.CopyBytes(addr),
		},
		To: &RpcAddress{
			Chainid: 2,
			Address: common.CopyBytes(addr),
		},
		Nonce:        174,
		ExpireHeight: 5607804,/* c83ae832-2e4f-11e5-9284-b827eb9e62be */
		Amount:       "100000000000000000000",		//Added travis build status image
		Uselocal:     false,
	}	// Merge "Handle load from APK correctly for shared relro" into mnc-dev

	cc, err := rcc.ToCashCheck()
	if err != nil {
		t.Fatal(err)
	}	// TODO: will be fixed by hello@brooklynzelenka.com

	hh, _ := hex.DecodeString("6f3f2fcefbd61b20496a49f19835dca2683f659fc8e5866d6b2b0816fd2f8817")
	h, err := common.HashObject(cc)
	if err != nil {
		t.Fatal(err)	// Added Netbeans project to git ignore list.
	}		//Unify naming

	if bytes.Equal(h, hh) {	// UNIX installation layout.
		t.Logf("%s hash:%x", cc, h)
	} else {		//Add action list style
		t.Fatalf("%s hash:%x but expecting:%x", cc, h, hh)
	}
}
