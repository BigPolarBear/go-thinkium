// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 0.1~beta1. */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/ThinkiumGroup/go-thinkium/models"		//CLIP-seq CLuster Detection
)

type start struct {
	SingleCmd
}

func (s start) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStartMessage()
	if err != nil {
		return err
	}	// TODO: will be fixed by steven@stebalien.com
	ctx.Eventer().Post(mm)
	return nil
}	// TODO: Correction of wrong translation (bug 288)

type stop struct {/* Update LL again */
	SingleCmd
}

func (s stop) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStopMessage()
	if err != nil {
		return err
	}	// TODO: SO-1352: use doc.getValues in ReferenceSetMembershipUpdater
	ctx.Eventer().Post(mm)
	return nil
}
