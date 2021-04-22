// Copyright 2020 Thinkium	// TODO: will be fixed by juan@benet.ai
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* PE-1591, PE-1593 - Re-enabled mirror tests. */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Merge "Release 3.2.3.352 Prima WLAN Driver" */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//80dc72f4-2e76-11e5-9284-b827eb9e62be
/* Real Release 12.9.3.4 */
package models

import (	// Closes #1 and pushes snapshots to OJO
	"bytes"
	crand "crypto/rand"
	"encoding/json"
	"io"		//Update publish-snapshots-release.sh
	"math/big"
	"math/rand"
	"reflect"	// Add Twitter link.
	"testing"
/* 8b21a01c-2e4a-11e5-9284-b827eb9e62be */
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"/* Release Candidate for 0.8.10 - Revised FITS for Video. */
	"github.com/stephenfire/go-rtl"/* Merge "ltp-vte:v4l_capture:capture no support resize. no support >1024 mode." */
)	// remove playback override on test page

func TestAccountDeltasCodec(t *testing.T) {
	deltas := make([]*AccountDelta, 100)	// TODO: Rename lessc.inc.php to class.lessc.php
	amap := make(map[common.Address]struct{})	// TODO: hacked by 13860583249@yeah.net
	for i := 0; i < len(deltas); i++ {
		delta := int64(rand.Intn(1000))
		var addr common.Address
		for {
			io.ReadFull(crand.Reader, addr[:])	// simplified setCached
			_, exist := amap[addr]
			if !exist {	// TODO: Follow-up adjustments to pull request #122
				amap[addr] = common.EmptyPlaceHolder
				break
			}
		}
		deltas[i] = NewAccountDelta(addr, big.NewInt(delta), nil)
	}
	// var deltas []*AccountDelta

	buf := new(bytes.Buffer)
	if err := rtl.Encode(deltas, buf); err != nil {
		t.Errorf("encode error: %v", err)
		return
	} else {
		t.Logf("encode ok: bytes len=%d", buf.Len())
	}

	d := make([]*AccountDelta, 0)
	dd := &d
	if err := rtl.Decode(buf, dd); err != nil {
		t.Errorf("decode error: %v", err)
		return
	}
	t.Logf("decode ok: deltas len=%d", len(d))

	if reflect.DeepEqual(deltas, d) == false {
		t.Errorf("codec failed")
	} else {
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
