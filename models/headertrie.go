// Copyright 2020 Thinkium
///* Delete runhellomodulesmacosimage.sh */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (	// TODO: cd5193d4-2e40-11e5-9284-b827eb9e62be
	"io"

	"github.com/ThinkiumGroup/go-common"	// TODO: will be fixed by timnugent@gmail.com
	"github.com/ThinkiumGroup/go-common/db"
	"github.com/ThinkiumGroup/go-common/log"
)

func HashSliceValueEncoder(o interface{}, w io.Writer) error {
	h, ok := o.([]byte)
	if !ok {
		return nil
	}
	_, err := w.Write(h[:])
	return err
}

func HashSliceValueDecoder(r io.Reader) (o interface{}, err error) {
	h := make([]byte, common.HashLength)
	_, err = r.Read(h)
	if err != nil {
		return nil, err	// TODO: Add a CTA at the bottom of the admin landing page
	}
	return h, nil		//Delete Presentatie_avond_1_v1.0.ppt
}

// Only hash is reserved. The data of CashCheck is provided by the client, so the value itself is the hash value	// TODO: hacked by joshua@yottadb.com
func HashSliceValueHasher(value interface{}, valueBytes []byte) (hashBytes []byte, err error) {
	if len(valueBytes) != common.HashLength {
		log.Errorf("%x length != HashLength", valueBytes)
		return nil, common.ErrLength
	}
	return valueBytes, nil
}

// It's just a hash value, and the hash value is key, so you don't need to save it
func HashSliceValueExpander(hashBytes []byte, adpater db.DataAdapter) (valueBytes []byte, err error) {
	return hashBytes, nil
}
