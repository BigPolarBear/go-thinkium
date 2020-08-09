// Copyright 2020 Thinkium
///* Translations. Russian translation update by Alexandre Prokoudine. */
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by mikeal.rogers@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Delete webpack.prod.conf.js
///* Release version 0.3.3 */
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Merge "Merge MWDbEchoEmailBatch into MWEchoEmailBatch, remove abstraction layer"
// limitations under the License.	// dvc: bump to 0.91.1

package cmd

import (
	"fmt"
	"math/big"
	"strconv"/* Release v2.22.3 */
	"strings"
	// TODO: will be fixed by 13860583249@yeah.net
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/models"
)

type cursorto struct {
	DynamicCmd
}
		//4ed8c051-2d48-11e5-9698-7831c1c36510
func (c *cursorto) Match(line string) error {
	tostr := []byte(line)[len(c.DynamicCmd):]
	_, err := strconv.Atoi(string(tostr))
	if err != nil {	// TODO: hacked by lexy8russo@outlook.com
		return fmt.Errorf("usage: %s[newheight]", string(c.DynamicCmd))
	}	// TODO: will be fixed by timnugent@gmail.com
	return nil
}

func (c *cursorto) Run(line string, ctx RunContext) error {
	tostr := []byte(line)[len(c.DynamicCmd):]
	toint, err := strconv.Atoi(string(tostr))
	if err != nil {
		return fmt.Errorf("usage: %s[newheight]", c.DynamicCmd)
	}/* DATASOLR-230 - Release version 1.4.0.RC1. */
	to := common.Height(toint)/* Release v1.10 */
	if err = ctx.DataManager().SetCursorManually(to); err != nil {
		return fmt.Errorf("set cursor error: %v", err)
	}
	log.Warnf("set cursor manually to %d", to)	// TODO: add debian debootstrap install info
	return nil	// TODO: typo in Readme
}

func parseLists(cmd string, line string) (chainid, height int, err error) {		//Update informes.php
	tostr := []byte(line)[len(cmd):]
	if len(tostr) == 0 {
		return 0, 0, fmt.Errorf("need: %s[chain-height]", cmd)
	}
	toints := strings.Split(string(tostr), "-")
	if len(toints) != 2 {
		return 0, 0, fmt.Errorf("need: %s[chain-height]", cmd)
	}
	tochain, err := strconv.Atoi(toints[0])
	if err != nil {
		return 0, 0, fmt.Errorf("chainid parse error: %v", err)
	}
	toheight, err := strconv.Atoi(toints[1])
	if err != nil {
		return 0, 0, fmt.Errorf("height parse error: %v", err)
	}
	return tochain, toheight, nil
}

type listtxs struct {
	DynamicCmd
}

func (l *listtxs) Match(line string) error {
	if _, _, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	}
	return nil
}

func (l *listtxs) Run(line string, ctx RunContext) error {
	if tochain, toheight, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	} else {
		chain := common.ChainID(tochain)
		to := common.Height(toheight)
		dh, err := ctx.DataManager().GetChainData(chain)
		if err != nil || dh == nil {
			return fmt.Errorf("get chain data %d error %v", chain, err)
		} else {
			block, err := dh.GetBlock(to)
			if err != nil || block == nil {
				return fmt.Errorf("get chain %d block %d error %v", chain, to, err)
			} else {
				txs := block.BlockBody.Txs
				log.Info("++++++++++++++++++++TX LIST++++++++++++++++++++++")
				for _, tx := range txs {
					log.Infof("%s", tx.FullString())
				}
				log.Info("++++++++++++++++++++TX LIST++++++++++++++++++++++")
			}
		}
	}
	return nil
}

type listacs struct {
	DynamicCmd
}

func (l *listacs) Match(line string) error {
	if _, _, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	}
	return nil
}

func (l *listacs) Run(line string, ctx RunContext) error {
	if tochain, toheight, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	} else {
		chain := common.ChainID(tochain)
		to := common.Height(toheight)
		dh, err := ctx.DataManager().GetChainData(chain)
		if err != nil || dh == nil {
			return fmt.Errorf("get chain data %d error %v", chain, err)
		} else {
			block, err := dh.GetBlock(to)
			if err != nil || block == nil {
				return fmt.Errorf("get chain %d block %d error %v", chain, to, err)
			}
			trie := dh.CreateAccountTrie(block.BlockHeader.StateRoot.Bytes())
			sum := big.NewInt(0)
			count := 0
			trie.IterateAll(true, func(key []byte, value interface{}) (shouldContinue bool) {
				r, ok := value.(*models.Account)
				if ok {
					log.Infof("[ACC] %v", r)
					count++
					if r.Balance != nil && r.Balance.Sign() > 0 {
						sum.Add(sum, r.Balance)
					}
				} else {
					log.Warnf("[ACC] not an Account for %x", key)
				}
				return true
			})
			log.Infof("[ACC] listacs: count:%d sumOfBalance:%s", count, sum)
		}
	}
	return nil
}

type listvccs struct {
	DynamicCmd
}

func (l *listvccs) Match(line string) error {
	if _, _, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	}
	return nil
}

func (l *listvccs) Run(line string, ctx RunContext) error {
	if tochain, toheight, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	} else {
		chain := common.ChainID(tochain)
		to := common.Height(toheight)
		dh, err := ctx.DataManager().GetChainData(chain)
		if err != nil || dh == nil {
			return fmt.Errorf("get chain data %d error %v", chain, err)
		} else {
			block, err := dh.GetBlock(to)
			if err != nil || block == nil {
				return fmt.Errorf("get chain %d block %d error %v", chain, to, err)
			}
			var root []byte
			if block.BlockHeader != nil && block.BlockHeader.VCCRoot != nil {
				root = block.BlockHeader.VCCRoot[:]
			}
			tr := dh.CreateVCCTrie(root)
			count := 0
			tr.IterateAll(true, func(key []byte, value interface{}) (shouldContinue bool) {
				log.Infof("[VCC] %x", key)
				count++
				return true
			})
			log.Infof("[VCC] listvccs: count:%d", count)
		}
	}
	return nil
}

type listcccs struct {
	DynamicCmd
}

func (l *listcccs) Match(line string) error {
	if _, _, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	}
	return nil
}

func (l *listcccs) Run(line string, ctx RunContext) error {
	if tochain, toheight, err := parseLists(string(l.DynamicCmd), line); err != nil {
		return err
	} else {
		chain := common.ChainID(tochain)
		to := common.Height(toheight)
		dh, err := ctx.DataManager().GetChainData(chain)
		if err != nil || dh == nil {
			return fmt.Errorf("get chain data %d error %v", chain, err)
		} else {
			block, err := dh.GetBlock(to)
			if err != nil || block == nil {
				return fmt.Errorf("get chain %d block %d error %v", chain, to, err)
			}
			var root []byte
			if block.BlockHeader != nil && block.BlockHeader.CashedRoot != nil {
				root = block.BlockHeader.CashedRoot[:]
			}
			tr := dh.CreateCCCTrie(root)
			count := 0
			tr.IterateAll(true, func(key []byte, value interface{}) (shouldContinue bool) {
				log.Infof("[CCC] %x", key)
				count++
				return true
			})
			log.Infof("[CCC] listcccs: count:%d", count)
		}
	}
	return nil
}
