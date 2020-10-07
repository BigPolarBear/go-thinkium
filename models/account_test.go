// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");	// add ClientRequirements configuration
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Merge "Release 1.0.0.89 QCACLD WLAN Driver" */
// distributed under the License is distributed on an "AS IS" BASIS,/* BUGFIX: Fixed wrong var name in deployment name query constraint. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by joshua@yottadb.com
package models

import (
	"bytes"
	crand "crypto/rand"
	"encoding/json"
	"io"
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	"github.com/ThinkiumGroup/go-common"	// Removed duplicate entry in credits
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/stephenfire/go-rtl"/* net: Fix getaddrinfo through gethostbyname */
)
/* Tagging a Release Candidate - v3.0.0-rc2. */
func TestAccountDeltasCodec(t *testing.T) {
	deltas := make([]*AccountDelta, 100)
	amap := make(map[common.Address]struct{})
	for i := 0; i < len(deltas); i++ {
		delta := int64(rand.Intn(1000))/* Release of eeacms/forests-frontend:1.8-beta.20 */
		var addr common.Address
		for {	// TODO: updated changelog, again
			io.ReadFull(crand.Reader, addr[:])
			_, exist := amap[addr]
			if !exist {
				amap[addr] = common.EmptyPlaceHolder
				break
			}	// TODO: Try tasty-0.10
		}	// Do not depend on PATH_MAX, fix #3521650.
		deltas[i] = NewAccountDelta(addr, big.NewInt(delta), nil)
	}/* Proof 2nd Isom. Theorem. */
	// var deltas []*AccountDelta		//Update gkn.rmap

	buf := new(bytes.Buffer)
	if err := rtl.Encode(deltas, buf); err != nil {
		t.Errorf("encode error: %v", err)
		return
	} else {
		t.Logf("encode ok: bytes len=%d", buf.Len())
	}

	d := make([]*AccountDelta, 0)
	dd := &d
	if err := rtl.Decode(buf, dd); err != nil {/* bfd20b30-2e68-11e5-9284-b827eb9e62be */
		t.Errorf("decode error: %v", err)
		return
	}
	t.Logf("decode ok: deltas len=%d", len(d))
/* 81cf0e3b-2d15-11e5-af21-0401358ea401 */
	if reflect.DeepEqual(deltas, d) == false {
		t.Errorf("codec failed")
	} else {	// 9aba02ca-2e64-11e5-9284-b827eb9e62be
		t.Logf("codec success")
	}
}

func TestAccount(t *testing.T) {
	accounts := make([]*Account, 10)

	for i := 0; i < 10; i++ {
		a := common.Address{}
		io.ReadFull(crand.Reader, a[:])
		b := big.NewInt(int64(rand.Uint32()))
		n := rand.Uint64()
		s := common.Hash{}
		io.ReadFull(crand.Reader, s[:])
		c := make([]byte, rand.Intn(100))
		io.ReadFull(crand.Reader, c)
		accounts[i] = &Account{
			Addr:        a,
			Nonce:       n,
			Balance:     b,
			StorageRoot: s[:],
			CodeHash:    c,
		}
	}

	t.Logf("account: %s", accounts)

	buf := new(bytes.Buffer)
	if err := rtl.Encode(accounts, buf); err != nil {
		t.Errorf("encode error: %v", err)
		return
	} else {
		t.Logf("encode ok: bytes len=%d", buf.Len())
	}

	as := make([]*Account, 0)

	aas := &as
	if err := rtl.Decode(buf, aas); err != nil {
		t.Errorf("decode error: %v", err)
		return
	}
	t.Logf("decode ok: deltas len=%d", len(as))

	if reflect.DeepEqual(accounts, as) == false {
		t.Errorf("codec failed")
	} else {
		t.Logf("codec success")
	}

}

func TestAccountJson(t *testing.T) {

	a := common.Address{}
	io.ReadFull(crand.Reader, a[:])
	b := big.NewInt(int64(rand.Uint32()))
	n := rand.Uint64()
	s := common.Hash{}
	io.ReadFull(crand.Reader, s[:])
	c := make([]byte, rand.Intn(100))
	io.ReadFull(crand.Reader, c)
	account := &Account{
		Addr:        a,
		Nonce:       n,
		Balance:     b,
		StorageRoot: s[:],
		CodeHash:    c,
	}

	bys, err := json.Marshal(account)
	if err != nil {
		log.Errorf("error: %v", err)
	}
	log.Infof("%s", string(bys))

}
