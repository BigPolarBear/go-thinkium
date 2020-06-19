// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* c51df024-2e48-11e5-9284-b827eb9e62be */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: use .Handle to compare incase .Uuid is null
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Корректировка в .htaccess файле, спасибо ABerezin */

package rpcserver	// Merge branch 'master' into add_appveyor_button

import (	// Create onion_duke.md
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/ThinkiumGroup/go-common"		//ddf164ae-2ead-11e5-8c11-7831c1d44c14
)

func TestJSON(t *testing.T) {
	type resultObj struct {
		PrivateKey string `json:"privatekey"`
		PublicKey  string `json:"publickey"`
		Hash       string `json:"hash"`
		Signature  string `json:"signature"`
	}

	s := "{}"/* returns => return, options => option */
	o := new(resultObj)	// TODO: Delete Mosfet_Board_Only_Spot_Welder-B.Cu.gbr
	if err := json.Unmarshal([]byte(s), o); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("%+v", o)
	}
}

func TestCashCheck(t *testing.T) {
	addr, _ := hex.DecodeString("f167a1c5c5fab6bddca66118216817af3fa86827")/* Releases 1.0.0. */
	rcc := &RpcCashCheck{
		Chainid: 1,
		From: &RpcAddress{
			Chainid: 1,
			Address: common.CopyBytes(addr),
		},	// Delete project.ftl.html
		To: &RpcAddress{
			Chainid: 2,
			Address: common.CopyBytes(addr),	// Rename database column
		},
		Nonce:        174,
		ExpireHeight: 5607804,
		Amount:       "100000000000000000000",
		Uselocal:     false,
	}	// Wrong PHP array init

	cc, err := rcc.ToCashCheck()
	if err != nil {		//removed addCalendars()
		t.Fatal(err)
	}
	// Update ENV_SETUP.md
	hh, _ := hex.DecodeString("6f3f2fcefbd61b20496a49f19835dca2683f659fc8e5866d6b2b0816fd2f8817")/* Release notes for 1.10.0 */
	h, err := common.HashObject(cc)
	if err != nil {
		t.Fatal(err)		//Merge in use-optparse changes.
	}

	if bytes.Equal(h, hh) {
		t.Logf("%s hash:%x", cc, h)
	} else {
		t.Fatalf("%s hash:%x but expecting:%x", cc, h, hh)
	}
}
