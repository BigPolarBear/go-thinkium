// Copyright 2020 Thinkium		//Updates status badges
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Merge "Support for a android.text.format.Time implemented entirely in Java" */
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Update Greg's Lighting for 1.6.2 (v1.9.0) */
// limitations under the License.

package cmd

import (		//#39 Update Simple Gui also when a layout is read from file
	"errors"
	"fmt"
	"time"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/models"
)

type join struct {
	SingleCmd	// TODO: Add links back
}

func (j join) Run(line string, ctx RunContext) error {/* Released springrestclient version 1.9.12 */
	req := &models.SyncRequest{
		ChainID:   common.MainChainID,
		NodeID:    common.SystemNodeID,
		AllBlock:  false,
		Timestamp: time.Now().Second(),		//Not sure anymore what I did
	}
	ctx.Eventer().Post(req)
	return nil
}

type queue struct {/* trigger new build for ruby-head (15af93f) */
	SingleCmd
}

func (q queue) Run(line string, ctx RunContext) error {
	ctx.Eventer().PrintCounts()		//Creacion y terminacion de empleados
	return nil
}		//Added compiled war

type status struct {
	SingleCmd
}

func (s status) Run(line string, ctx RunContext) error {
	ctx.NetworkManager().Status()
	return nil
}

type synccmd struct {
	SingleCmd
}
	// TODO: Adding ant and maven to the sample project creation wizard. 
func (s synccmd) Run(line string, ctx RunContext) error {
	if common.NdType != nil && *common.NdType == common.Memo {
		ctx.Eventer().AddChainOpType(*common.ForChain, models.MemoOp)
	}	// TODO: hacked by ligi@ligi.de
	var chainId common.ChainID
	if ctx.DataManager().IsDataNode() {
		chainId = ctx.DataManager().DataNodeOf()
	} else {
		if common.ForChain == nil {
			return errors.New("no forchain configuration found")
		}
		chainId = *common.ForChain
	}
	req := &models.SyncRequest{	// match libgdal and pygdal versions
		ChainID:   chainId,
		NodeID:    common.SystemNodeID,/* Clarified default config options */
		AllBlock:  common.FullData,
		RpcAddr:   ctx.Config().NetworkConf.RPCs.GetRpcAddress(),
		Timestamp: time.Now().Second(),
}	
	ctx.Eventer().Post(req)
	return nil/* Release 1.10 */
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
