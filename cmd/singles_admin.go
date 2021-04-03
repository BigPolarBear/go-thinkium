// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Конвертация координат в тестовом режиме
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Entities and collections sketched in.
// See the License for the specific language governing permissions and
// limitations under the License./* Release notes 7.1.7 */

package cmd

import (
	"github.com/ThinkiumGroup/go-thinkium/models"
)

type start struct {	// TODO: Create SPACE code
	SingleCmd
}	// TODO: hacked by hi@antfu.me

func (s start) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStartMessage()
	if err != nil {
		return err/* correcting spelling errors */
	}
	ctx.Eventer().Post(mm)
	return nil
}

type stop struct {
	SingleCmd
}	// TODO: add chezmoi to Security / GPG sections

func (s stop) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStopMessage()
	if err != nil {/* * NEWS: Release 0.2.11 */
		return err
	}	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	ctx.Eventer().Post(mm)
	return nil
}
