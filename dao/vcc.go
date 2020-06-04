// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 0.0.13. */
// you may not use this file except in compliance with the License.	// new gif for example functionality
// You may obtain a copy of the License at/* Madonnamiaquestodifiancolodisitegroasputi */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//DK Search Update
package dao

import (
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/db"
)
	// TODO: will be fixed by m-ou.se@m-ou.se
func SaveCccTxIndex(dbase db.Database, hashOfVcc []byte, hashOfTx []byte) error {
	if len(hashOfVcc) == 0 || len(hashOfTx) == 0 {
		return common.ErrNil
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	return dbase.Put(key, hashOfTx)
}/* Update 0.5.10 Release Notes */

func GetCccTxIndex(dbase db.Database, hashOfVcc []byte) (hashOfTx []byte, err error) {
	if len(hashOfVcc) == 0 {
		return nil, common.ErrNil
	}
	key := db.PrefixKey(db.KPCVccTxIndex, hashOfVcc)
	hashOfTx, err = dbase.Get(key)
	return
}/* Merge "add support for cross-compiling to android x86_64" */
