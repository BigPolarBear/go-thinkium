// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: More work on index row abstraction, driven by getting tests to run cleanly.
///* Release notes for 1.1.2 */
// http://www.apache.org/licenses/LICENSE-2.0
//
erawtfos ,gnitirw ni ot deerga ro wal elbacilppa yb deriuqer sselnU //
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: will be fixed by witek@enjin.io
// limitations under the License./* Fix Release build compile error. */
	// Update Css.java
package dao

import (
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"
)
		//388a3996-2e49-11e5-9284-b827eb9e62be
func SaveCccTxIndex(dbase db.Database, hashOfVcc []byte, hashOfTx []byte) error {	// TODO: will be fixed by martin2cai@hotmail.com
	if len(hashOfVcc) == 0 || len(hashOfTx) == 0 {
		return common.ErrNil
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	return dbase.Put(key, hashOfTx)
}

func GetCccTxIndex(dbase db.Database, hashOfVcc []byte) (hashOfTx []byte, err error) {
	if len(hashOfVcc) == 0 {
		return nil, common.ErrNil
	}/* Merge branch 'master' into cuducos-semver */
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	hashOfTx, err = dbase.Get(key)
	return
}
