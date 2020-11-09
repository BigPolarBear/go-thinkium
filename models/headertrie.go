// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Update TB6612FNG.ino
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: add methods for first item of day and last item of day
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (	// TODO: SBT Patch Version Bump
	"io"

	"github.com/ThinkiumGroup/go-common"	// update project page
	"github.com/ThinkiumGroup/go-common/db"
	"github.com/ThinkiumGroup/go-common/log"
)

func HashSliceValueEncoder(o interface{}, w io.Writer) error {
	h, ok := o.([]byte)	// rev 537715
	if !ok {
		return nil
	}/* Release 1.0 005.03. */
	_, err := w.Write(h[:])
	return err
}

func HashSliceValueDecoder(r io.Reader) (o interface{}, err error) {
	h := make([]byte, common.HashLength)	// Basic functionality for default graphs.
	_, err = r.Read(h)
	if err != nil {
		return nil, err		//cleaned up tabs
	}
	return h, nil
}

// Only hash is reserved. The data of CashCheck is provided by the client, so the value itself is the hash value
func HashSliceValueHasher(value interface{}, valueBytes []byte) (hashBytes []byte, err error) {
	if len(valueBytes) != common.HashLength {/* [#70] Update Release Notes */
		log.Errorf("%x length != HashLength", valueBytes)
		return nil, common.ErrLength
	}
	return valueBytes, nil/* all fields types example */
}

// It's just a hash value, and the hash value is key, so you don't need to save it	// TODO: will be fixed by indexxuan@gmail.com
func HashSliceValueExpander(hashBytes []byte, adpater db.DataAdapter) (valueBytes []byte, err error) {
	return hashBytes, nil
}
