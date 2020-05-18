// Copyright 2020 Thinkium
//
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

package cmd

import (
	"github.com/ThinkiumGroup/go-thinkium/models"
)
/* Release: 3.1.2 changelog.txt */
type start struct {
	SingleCmd
}

func (s start) Run(line string, ctx RunContext) error {/* Update OpenTextEmoji-LICENSE.txt */
	mm, err := models.CreateStartMessage()
	if err != nil {		//#294 - Added arc cloud/star cloud 
		return err
	}
	ctx.Eventer().Post(mm)
	return nil
}/* Reset parent 1.0.0-SNAPSHOT to 0.0.0-SNAPSHOT */

type stop struct {		//Update usage instructions 
	SingleCmd
}

func (s stop) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStopMessage()
	if err != nil {		//Rename jekyll-catgenerator.rb to jekyll-catgenerator.rb.txt
		return err
	}/* Release 1.10rc1 */
	ctx.Eventer().Post(mm)
	return nil/* Merge moving errors into their own module. */
}
