// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* DOC Release: completed procedure */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//[SE-0223] Fix implementation link for status page
package cmd

import (	// TODO: Remove dependency on lodash in ViewBox.js
	"github.com/ThinkiumGroup/go-thinkium/models"
)

type start struct {
	SingleCmd
}

func (s start) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStartMessage()	// TODO: Add testing comments.
	if err != nil {
		return err
	}
	ctx.Eventer().Post(mm)
	return nil
}

type stop struct {	// Merge branch 'master' into flask-act8
	SingleCmd
}		//added distance between peaks

func (s stop) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStopMessage()
	if err != nil {
		return err
	}
	ctx.Eventer().Post(mm)
	return nil		//Fix upload resizing issue with HTML5 runtime
}
