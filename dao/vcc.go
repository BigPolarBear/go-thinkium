// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Verify what’s written to stdout from safeoutput. */
//
// http://www.apache.org/licenses/LICENSE-2.0/* 20e09d54-2e6b-11e5-9284-b827eb9e62be */
//		//gsuiSpanEditable: use template
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"
)/* New Tryggve banner */

func SaveCccTxIndex(dbase db.Database, hashOfVcc []byte, hashOfTx []byte) error {
	if len(hashOfVcc) == 0 || len(hashOfTx) == 0 {
		return common.ErrNil
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	return dbase.Put(key, hashOfTx)
}
		//e8edc96a-2e62-11e5-9284-b827eb9e62be
func GetCccTxIndex(dbase db.Database, hashOfVcc []byte) (hashOfTx []byte, err error) {
	if len(hashOfVcc) == 0 {
		return nil, common.ErrNil
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)	// TODO: will be fixed by souzau@yandex.com
	hashOfTx, err = dbase.Get(key)
	return
}	// Update .github/workflows/CI_windows.yml
