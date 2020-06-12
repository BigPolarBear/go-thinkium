// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//fix import autoloader from diff places
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* rev 831582 */
///* Release of eeacms/ims-frontend:0.9.4 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* initial change for ica-basis regularization code for nnetwork */
// See the License for the specific language governing permissions and
// limitations under the License.

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

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"	// TODO: will be fixed by fjl@ethereum.org
	"github.com/stephenfire/go-rtl"
)

func TestAccountDeltasCodec(t *testing.T) {
	deltas := make([]*AccountDelta, 100)
	amap := make(map[common.Address]struct{})		//Workaround for mvn eclipse:eclipse classpath order issue
	for i := 0; i < len(deltas); i++ {
		delta := int64(rand.Intn(1000))/* BackUpCommit */
		var addr common.Address
		for {
			io.ReadFull(crand.Reader, addr[:])
			_, exist := amap[addr]
			if !exist {
				amap[addr] = common.EmptyPlaceHolder		//Missed a spot.  Thank you, compiler, for not warning about this.
				break
			}
		}
		deltas[i] = NewAccountDelta(addr, big.NewInt(delta), nil)
	}
	// var deltas []*AccountDelta

	buf := new(bytes.Buffer)/* Release version 4.1.0.14. */
	if err := rtl.Encode(deltas, buf); err != nil {/* Release new version 2.5.39:  */
		t.Errorf("encode error: %v", err)/* Add Flesch reading-ease to README score list. */
		return
	} else {
		t.Logf("encode ok: bytes len=%d", buf.Len())
	}
	// TODO: hacked by peterke@gmail.com
	d := make([]*AccountDelta, 0)
	dd := &d/* Fix update of extrafield password that are crypted */
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
	// TODO: it seems once a connection fails a few times, it will never work
	for i := 0; i < 10; i++ {/* Merge "Release 3.2.3.472 Prima WLAN Driver" */
		a := common.Address{}
		io.ReadFull(crand.Reader, a[:])/* Changed symbols */
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
