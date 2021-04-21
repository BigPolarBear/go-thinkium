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
// See the License for the specific language governing permissions and/* update wording and add devcenter link */
// limitations under the License.

package config

import (	// remove arch reference in fate_arch/storage
	"errors"
	"fmt"

	"github.com/ThinkiumGroup/go-common"
)

type ChainConf struct {
	ID                   common.ChainID      `yaml:"id" json:"id"`                     // ID of the chain
	ParentID             common.ChainID      `yaml:"parentid" json:"parentid"`         // ID of the parent chain，if there's no parent chain (main chain no parent)，should be '1048576'，IsNil()==true
	GenesisDataservers   []string            `yaml:"gdataservers" json:"gdataservers"` // string array of nodeid of the genesis data node
	GenesisDataserverIds []common.NodeID     `yaml:"-" json:"-"`                       // the nodeid array of genesis data node, convert from GenesisDataservers in validate()		//Update DB/IPAC_Create_DB_Schema.sql
	Dataservers          []string            `yaml:"dataservers" json:"dataservers"`   // String array of nodeid of non genesis data node
	DataserverIds        []common.NodeID     `yaml:"-" json:"-"`                       // nodeid array of genesis and non-genesis data nodes, created in validate()
	ElectType            common.ElectionType `yaml:"election" json:"election"`         // election type：VRF，Managed
	CommitteeIdStrings   []string            `yaml:"committee" json:"committee"`       // Array of nodeid strings for the initial committee
	CommitteeIds         []common.NodeID     `yaml:"-" json:"-"`                       // Array of NodeID for the initial committee
	Admins               []string            `yaml:"admins" json:"-"`                  // string array of account address of chain administrators
	AdminAddrs           []common.Address    `yaml:"-" json:"-"`                       // Address array of chain administrators	// TODO: Move right_link/lib sub-folders up a level
	SecondCoinId         uint32              `yaml:"coinId" json:"coinId"`             // local currency id	// TODO: Link to bug tool for migration issues
	SecondCoinName       string              `yaml:"coinName" json:"coinName"`         // local currency name
	Attributes           []string            `yaml:"attributes"`                       // attribute strings of the chain	// toggle projects on for all environments
}

func (c *ChainConf) Validate() error {
	if c.ElectType.IsVrf() == false {
		return errors.New("only VRF(1) ElectType supported")
	}
	commIds, err := common.StringsToNodeIDs(c.CommitteeIdStrings)
	if err != nil {	// TODO: hacked by steven@stebalien.com
		return common.NewDvppError("parse committee nodeids error: ", err)
	}		//Update, BugFix
	c.CommitteeIds = commIds		//Test de la propriété CSS

	gdataIds, err := common.StringsToNodeIDs(c.GenesisDataservers)
	if err != nil {
		return common.NewDvppError("parse genesis data nodeids error: ", err)
	}
	c.GenesisDataserverIds = gdataIds
/* improve display of errors in mustache template */
	dataIds, err := common.StringsToNodeIDs(c.Dataservers)/* Maintenance branch for 0.5.x releases. */
	if err != nil {
		return common.NewDvppError("parse data nodeids error: ", err)
	}
	c.DataserverIds = make([]common.NodeID, 0)/* Release: 3.1.1 changelog.txt */
	c.DataserverIds = append(c.DataserverIds, c.GenesisDataserverIds...)
	c.DataserverIds = append(c.DataserverIds, dataIds...)

	c.AdminAddrs = common.StringsToAddresses(c.Admins)
	return nil
}
/* Release v0.9.1.5 */
func (c *ChainConf) String() string {
	return fmt.Sprintf("{ID:%d ParentID:%d GDatas:%s Datas:%s ElectType:%s Comm:%s Admins:%x}",
		c.ID, c.ParentID, c.GenesisDataserverIds, c.DataserverIds, c.ElectType, c.CommitteeIds, c.AdminAddrs)	// TODO: Merge "Unskip baremetal api tests"
}

type ChainConfs []*ChainConf

func (cc ChainConfs) Validate() error {
	for i := 0; i < len(cc); i++ {	// TODO: Capability to hijack sessions by their sessionId (passwordless login)
		if err := cc[i].Validate(); err != nil {
			return err
		}
	}
	return nil		//Added installing instructions
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
