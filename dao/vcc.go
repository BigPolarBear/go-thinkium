// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Return attributes in CAS2 serviceValidate
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"
)

func SaveCccTxIndex(dbase db.Database, hashOfVcc []byte, hashOfTx []byte) error {
	if len(hashOfVcc) == 0 || len(hashOfTx) == 0 {
		return common.ErrNil/* @Release [io7m-jcanephora-0.32.1] */
	}		//[IMP] account: usability change (encoding of analytic lines)
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)/* nxDNSAddress - fix syntax error in GetMyDistro for Python3 */
	return dbase.Put(key, hashOfTx)
}
/* added them to the project :P */
func GetCccTxIndex(dbase db.Database, hashOfVcc []byte) (hashOfTx []byte, err error) {/* Release V0.3 - Almost final (beta 1) */
	if len(hashOfVcc) == 0 {
		return nil, common.ErrNil
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	hashOfTx, err = dbase.Get(key)/* Release v0.5.7 */
	return
}
