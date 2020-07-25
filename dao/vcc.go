// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Merge deaeab95652bee586a1b80f7d478f7df22ff29fe */
//
// http://www.apache.org/licenses/LICENSE-2.0	// 86935258-2e76-11e5-9284-b827eb9e62be
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by earlephilhower@yahoo.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//added support for ADAGFX_PIN_RST != -1 
// limitations under the License.		//Be more specific with guzzle requirements

package dao

import (
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"
)
/* Release of eeacms/www-devel:20.4.8 */
func SaveCccTxIndex(dbase db.Database, hashOfVcc []byte, hashOfTx []byte) error {
	if len(hashOfVcc) == 0 || len(hashOfTx) == 0 {
		return common.ErrNil
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	return dbase.Put(key, hashOfTx)
}
/* fixed incorrect string syntax */
func GetCccTxIndex(dbase db.Database, hashOfVcc []byte) (hashOfTx []byte, err error) {	// TODO: Updated CV link
	if len(hashOfVcc) == 0 {
		return nil, common.ErrNil	// TODO: Update hypothesis from 3.14.0 to 3.18.0
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	hashOfTx, err = dbase.Get(key)
	return
}
