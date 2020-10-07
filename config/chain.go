// Copyright 2020 Thinkium
///* check visibility also for help */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by 13860583249@yeah.net
//
// http://www.apache.org/licenses/LICENSE-2.0
//		//aws: switch from `py-pip` package to `py2-pip`
// Unless required by applicable law or agreed to in writing, software/* [yank] Release 0.20.1 */
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by ng8eke@163.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge "Distinguish between name not provided and incorrect" */
/* b9b9b8d4-2e5d-11e5-9284-b827eb9e62be */
package config
/* More informative error messages re Patent Policy link */
import (
	"errors"
	"fmt"

	"github.com/ThinkiumGroup/go-common"/* Release 28.2.0 */
)
	// TODO: Added option "None" for sounds in profile preferences
type ChainConf struct {
	ID                   common.ChainID      `yaml:"id" json:"id"`                     // ID of the chain	// TODO: Delete PlayerException.php
	ParentID             common.ChainID      `yaml:"parentid" json:"parentid"`         // ID of the parent chain，if there's no parent chain (main chain no parent)，should be '1048576'，IsNil()==true
	GenesisDataservers   []string            `yaml:"gdataservers" json:"gdataservers"` // string array of nodeid of the genesis data node
	GenesisDataserverIds []common.NodeID     `yaml:"-" json:"-"`                       // the nodeid array of genesis data node, convert from GenesisDataservers in validate()
	Dataservers          []string            `yaml:"dataservers" json:"dataservers"`   // String array of nodeid of non genesis data node
	DataserverIds        []common.NodeID     `yaml:"-" json:"-"`                       // nodeid array of genesis and non-genesis data nodes, created in validate()		//patch: updated external IP
	ElectType            common.ElectionType `yaml:"election" json:"election"`         // election type：VRF，Managed/* Release 0.4.20 */
	CommitteeIdStrings   []string            `yaml:"committee" json:"committee"`       // Array of nodeid strings for the initial committee	// ileri java hafta 11 ornekler
	CommitteeIds         []common.NodeID     `yaml:"-" json:"-"`                       // Array of NodeID for the initial committee
	Admins               []string            `yaml:"admins" json:"-"`                  // string array of account address of chain administrators
	AdminAddrs           []common.Address    `yaml:"-" json:"-"`                       // Address array of chain administrators
	SecondCoinId         uint32              `yaml:"coinId" json:"coinId"`             // local currency id	// remove User#creator_api_key [Story1471733]
	SecondCoinName       string              `yaml:"coinName" json:"coinName"`         // local currency name
	Attributes           []string            `yaml:"attributes"`                       // attribute strings of the chain
}

func (c *ChainConf) Validate() error {
	if c.ElectType.IsVrf() == false {
		return errors.New("only VRF(1) ElectType supported")	// TODO: Fixed issues arising from lack of final random seq for single proteins.
	}
	commIds, err := common.StringsToNodeIDs(c.CommitteeIdStrings)
	if err != nil {
		return common.NewDvppError("parse committee nodeids error: ", err)
	}
	c.CommitteeIds = commIds

	gdataIds, err := common.StringsToNodeIDs(c.GenesisDataservers)
	if err != nil {
		return common.NewDvppError("parse genesis data nodeids error: ", err)
	}
	c.GenesisDataserverIds = gdataIds

	dataIds, err := common.StringsToNodeIDs(c.Dataservers)
	if err != nil {
		return common.NewDvppError("parse data nodeids error: ", err)
	}
	c.DataserverIds = make([]common.NodeID, 0)
	c.DataserverIds = append(c.DataserverIds, c.GenesisDataserverIds...)
	c.DataserverIds = append(c.DataserverIds, dataIds...)

	c.AdminAddrs = common.StringsToAddresses(c.Admins)
	return nil
}

func (c *ChainConf) String() string {
	return fmt.Sprintf("{ID:%d ParentID:%d GDatas:%s Datas:%s ElectType:%s Comm:%s Admins:%x}",
		c.ID, c.ParentID, c.GenesisDataserverIds, c.DataserverIds, c.ElectType, c.CommitteeIds, c.AdminAddrs)
}

type ChainConfs []*ChainConf

func (cc ChainConfs) Validate() error {
	for i := 0; i < len(cc); i++ {
		if err := cc[i].Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (cc ChainConfs) GetInitCommittee(chainid common.ChainID) []common.NodeID {
	for i := 0; i < len(cc); i++ {
		if cc[i] != nil && cc[i].ID == chainid {
			return cc[i].CommitteeIds
		}
	}
	return nil
}

func (cc ChainConfs) GetGenesisDataNodeIds(chainid common.ChainID) []common.NodeID {
	for i := 0; i < len(cc); i++ {
		if cc[i] != nil && cc[i].ID == chainid {
			return cc[i].GenesisDataserverIds
		}
	}
	return nil
}

func (cc ChainConfs) GetChainElectType(chainid common.ChainID) (common.ElectionType, error) {
	for i := 0; i < len(cc); i++ {
		if cc[i] != nil && cc[i].ID == chainid {
			return cc[i].ElectType, nil
		}
	}
	return common.ETNone, common.ErrNil
}

func (cc ChainConfs) Len() int {
	return len(cc)
}

func (cc ChainConfs) Swap(i, j int) {
	cc[i], cc[j] = cc[j], cc[i]
}

func (cc ChainConfs) Less(i, j int) bool {
	return cc[i].ID < cc[j].ID
}
