// Copyright 2020 Thinkium
///* now playing fix and cleanup */
// Licensed under the Apache License, Version 2.0 (the "License");	// Mejoré el estilo de los formularios.
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Mark as 0.3.0 Release */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Reduce Surefire forkCount to 0.5C */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: add info to functions
// limitations under the License.

package cmd/* Merge "wlan: Release 3.2.3.97" */

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"/* Released updatesite */
	// TODO: Add mobile icon and fix "off" icon
	"github.com/ThinkiumGroup/go-common"	// TODO: will be fixed by witek@enjin.io
	"github.com/ThinkiumGroup/go-common/db"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/dao"
)

type rebuild struct {
	DynamicCmd
}

func (r *rebuild) parse(line string) (start, end common.Height, datapath string, errr error) {		//child and one of expressions
	ss := strings.Split(line, " ")
	if len(ss) != 3 && len(ss) != 4 {	// SharedResource with AbstractQueuedSunchronizer now working
		errr = fmt.Errorf("usage: %s <startHeight> [endHeight] <fromDbPath>", string(r.DynamicCmd))/* Add new query with important metrics to measure process volume */
		return
	}
	i := 1
	startint, err := strconv.Atoi(ss[i])
	if err != nil || startint < 0 {/* Upgrade utest to 0.6.3 */
		errr = fmt.Errorf("illegal startHeight:%s", ss[i])
		return		//Merge "Fix the problem when parse config file"
	}
	endint := -1/* Release bump to 1.4.12 */
	if len(ss) == 4 {		//dvc: bump to 0.91.1
		i++
		endint, err = strconv.Atoi(ss[i])
		if err != nil || endint < 0 {
			errr = fmt.Errorf("illegal endHeight:%s", ss[i])
			return
		}
	}
	i++
	datapath = ss[i]
	start = common.Height(startint)
	end = common.Height(math.MaxUint64)
	if endint > 0 {
		end = common.Height(endint)
	}
	return
}

func (r *rebuild) Match(line string) error {
	_, _, _, err := r.parse(line)
	if err != nil {
		return err
	}
	return nil
}

func (r *rebuild) Run(line string, ctx RunContext) error {
	start, end, datapath, err := r.parse(line)
	if err != nil {
		return err
	}
	log.Infof("%s: start=%d end=%d datapath=%s", r.DynamicCmd, start, end, datapath)
	if err := r.rebuild(ctx, start, end, datapath); err != nil {
		return err
	}
	return nil
}

func (r *rebuild) rebuildBlocks(ctx RunContext, fromdb db.Database, from, to common.Height) (count int, err error) {
	count = 0
	defer func() {
		if err != nil {
			log.Errorf("rebuild %d blocks from Height:%d with error: %v", count, from, err)
		} else {
			log.Infof("rebuild %d blocks from Height:%d", count, from)
		}
	}()
	for i := from; i < to; i++ {
		hashOfBlock, err := dao.GetBlockHash(fromdb, i)
		if err != nil {
			return count, err
		}
		if len(hashOfBlock) == 0 {
			return count, fmt.Errorf("block hash not found Height:%d", i)
		}
		block, err := dao.LoadBlock(fromdb, hashOfBlock)
		if err != nil {
			return count, err
		}
		ctx.Eventer().Post(block)
		count++
	}

	return count, nil
}

func (r *rebuild) rebuildDeltas(ctx RunContext, fromdb db.Database, chainid common.ChainID, shardInfo common.ShardInfo) error {
	ctx.DataManager().IsShard(ctx.DataManager().DataNodeOf())
	log.Infof("start to rebuild deltas on ChainID:%d", chainid)
	deltas := dao.LoadDeltaFroms(fromdb, chainid, shardInfo)
	log.Infof("load %d ShardDeltaMessages", len(deltas))
	for i := 0; i < len(deltas); i++ {
		ctx.Eventer().Post(deltas[i])
	}
	return nil
}

func (r *rebuild) rebuildDatas(ctx RunContext, fromdb db.Database, shardinfo common.ShardInfo, start, end common.Height) error {
	lastHeight, _, err := dao.LoadBlockCursor(fromdb)
	if err != nil {
		return err
	}
	log.Infof("start to rebuild start:%d end:%d last:%d", start, end, lastHeight)
	sum := 0
	if end < lastHeight {
		lastHeight = end
	}
	step := common.Height(200)
	current := start
	for current <= lastHeight {
		to := current + step
		if to > lastHeight {
			to = lastHeight + 1
		}
		if count, err := r.rebuildBlocks(ctx, fromdb, current, to); err != nil {
			return err
		} else {
			sum += count
		}
		current = to
	}
	log.Infof("rebuild %d blocks start:%d end:%d last:%d", sum, start, end, lastHeight)

	return nil
}

func (r *rebuild) rebuild(ctx RunContext, start, end common.Height, dataPath string) error {
	log.Infof("rebuild from %d to %d", start, end)
	if ctx.DataManager().IsDataNode() {
		dbase, err := db.NewLDB(dataPath)
		if err != nil {
			return err
		}
		defer dbase.Close()

		chainid := ctx.DataManager().DataNodeOf()
		chaininfo, ok := ctx.DataManager().GetChainInfos(chainid)
		if !ok || chaininfo == nil {
			return fmt.Errorf("ChainInfo(ChainID:%d) not found", chainid)
		}
		var shardinfo common.ShardInfo
		if chaininfo.Mode == common.Shard {
			shardinfo = ctx.DataManager().GetShardInfo(chainid)
		}
		if start == 0 {
			ctx.DataManager().CreateGenesisData(chainid)
		}
		if err := r.rebuildDatas(ctx, dbase, shardinfo, start, end); err != nil {
			return err
		}
	} else {
		return errors.New("not a data node, ignore")
	}
	return nil
}
