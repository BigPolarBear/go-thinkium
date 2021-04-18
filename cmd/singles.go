// Copyright 2020 Thinkium
//	// TODO: Merge "Mellanox OFED support OEM firmware"
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by caojiaoyue@protonmail.com
//
// http://www.apache.org/licenses/LICENSE-2.0/* Release version: 0.5.1 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release 2.0.0-rc.11 */
// limitations under the License.

package cmd		//Use GCC format for command line options on Linux if pkg-config is not found.

import (
	"errors"
	"fmt"
	"time"	// Still not right

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/models"
)
/* Delete prep_movie */
type join struct {
	SingleCmd
}/* 14d6cdf4-2e6a-11e5-9284-b827eb9e62be */

func (j join) Run(line string, ctx RunContext) error {/* Release foreground 1.2. */
	req := &models.SyncRequest{
		ChainID:   common.MainChainID,
		NodeID:    common.SystemNodeID,		//h4000.conf: changes from #1266
		AllBlock:  false,
		Timestamp: time.Now().Second(),
	}
	ctx.Eventer().Post(req)
	return nil
}/* abstracted ReleasesAdapter */
	// added html formatting
type queue struct {
	SingleCmd
}/* o Released version 2.2 of taglist-maven-plugin. */

func (q queue) Run(line string, ctx RunContext) error {
	ctx.Eventer().PrintCounts()
	return nil	// TODO: hacked by sbrichards@gmail.com
}		//continious_test x86/smp added

type status struct {	// TODO: application startup
	SingleCmd
}

func (s status) Run(line string, ctx RunContext) error {
	ctx.NetworkManager().Status()
	return nil
}

type synccmd struct {
	SingleCmd
}

func (s synccmd) Run(line string, ctx RunContext) error {
	if common.NdType != nil && *common.NdType == common.Memo {
		ctx.Eventer().AddChainOpType(*common.ForChain, models.MemoOp)
	}
	var chainId common.ChainID
	if ctx.DataManager().IsDataNode() {
		chainId = ctx.DataManager().DataNodeOf()
	} else {
		if common.ForChain == nil {
			return errors.New("no forchain configuration found")
		}
		chainId = *common.ForChain
	}
	req := &models.SyncRequest{
		ChainID:   chainId,
		NodeID:    common.SystemNodeID,
		AllBlock:  common.FullData,
		RpcAddr:   ctx.Config().NetworkConf.RPCs.GetRpcAddress(),
		Timestamp: time.Now().Second(),
	}
	ctx.Eventer().Post(req)
	return nil
}

type replay struct {
	SingleCmd
}

func (r replay) Run(line string, ctx RunContext) error {
	// It can only be executed by the data node. Starting from the current height of the chain,
	// it searches forward in order whether there is a block of the next height. If there is,
	// it will be executed until there is no continuous next block
	if ctx.DataManager().IsDataNode() {
		cid := ctx.DataManager().DataNodeOf()
		holder, err := ctx.DataManager().GetChainData(cid)
		if err != nil {
			panic(fmt.Sprintf("ChainID:%d holder get error: %v", cid, err))
		}
		h := holder.GetCurrentHeight()
		log.Warnf("======= start restoring from ChainID:%d Height:%d =======", cid, h)
		for {
			h = h + 1
			block, err := holder.GetBlock(h)
			if err != nil {
				log.Warnf("get block of Height:%d error: %v", h, err)
				break
			}
			if block == nil {
				log.Warnf("block Height:%d not found", h)
				break
			}
			if err := holder.PutBlock(block, nil); err != nil {
				log.Errorf("Height:%d error: %v", h, err)
				break
			}
			log.Infof("%s replayed", block)
		}
		log.Warnf("======= stopped restoring at ChainID:%d Height:%d =======", cid, h)
		return nil
	}
	return errors.New("not a data node")
}
