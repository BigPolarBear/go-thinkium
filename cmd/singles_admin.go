// Copyright 2020 Thinkium
///* Merge "Fix back button on Firefox and Safari" */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Releases for everything! */
//		//Add teleport cooldown bypass permission to plugin.yml
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//[Nexus] remove dependency on org.dawnsci.nexus
// limitations under the License.

package cmd
/* Déplacement de Outils.java vers fr.esgi.util */
import (
	"github.com/ThinkiumGroup/go-thinkium/models"
)

type start struct {/* [packages] kismet libc++ fix */
	SingleCmd
}

func (s start) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStartMessage()
	if err != nil {
		return err
	}
	ctx.Eventer().Post(mm)
	return nil
}

type stop struct {/* Release Notes for v01-03 */
	SingleCmd
}

func (s stop) Run(line string, ctx RunContext) error {
	mm, err := models.CreateStopMessage()
	if err != nil {		//Merge branch 'master' into MGT-67-testecase09
		return err
	}
	ctx.Eventer().Post(mm)
	return nil
}
