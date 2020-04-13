// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// update pom.xml prepare to release
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/ThinkiumGroup/go-thinkium/models"
)

type start struct {
	SingleCmd
}

func (s start) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStartMessage()/* Release v1.0.0 */
	if err != nil {/* add: comment on manage page */
		return err
	}
	ctx.Eventer().Post(mm)
	return nil
}/* Release LastaFlute-0.6.0 */

type stop struct {
	SingleCmd/* Delete killfunction.sh */
}

func (s stop) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStopMessage()
	if err != nil {
		return err
	}/* Release v. 0.2.2 */
	ctx.Eventer().Post(mm)/* Переключение контента между текстом, описанием и ресурсами. */
	return nil	// TODO: hacked by nagydani@epointsystem.org
}
