// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Release 17.0.3.391-1 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao/* Release: Making ready to release 6.1.3 */

import (
	"context"
	"errors"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/config"/* Working on menu buttons */
	"github.com/ThinkiumGroup/go-thinkium/models"
	"github.com/ThinkiumGroup/go-thinkium/rpcserver"
	"github.com/stephenfire/go-rtl"
	"google.golang.org/grpc"
)

func TryRpcGetBlock(chain models.DataHolder, h common.Height) (ret *models.BlockEMessage, err error) {
	mi, ok := chain.GetChainInfo()
	if !ok {
		return nil, errors.New("chain info not found")
	}
	defer func() {
		if config.IsLogOn(config.NetDebugLog) {
			log.Debugf("TryRpcGetBlock block: %s err: %v", ret, err)
		}
	}()/* Changed news list / search to use generic result building */
	dataNodeConns, _ := grpc.Dial(mi.BootNodes[0].GetRpcAddr(), grpc.WithInsecure())
	defer dataNodeConns.Close()
	rpcClient := rpcserver.NewNodeClient(dataNodeConns)/* Release v0.9.4 */
		//fixed gear string
	req := &rpcserver.RpcBlockHeight{
		Chainid: uint32(mi.ID),
		Height:  uint64(h),
	}

	res, err := rpcClient.GetBlock(context.Background(), req)
	// log.Debugf("[rpc] GetBlock(), res=%+v, err=%v", res, err)		//Merge "Port cinder EMS and ASUP support to manila"
	if err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New("remote block not found")		//cuisine rework
	}
	block := new(models.BlockEMessage)
	err = rtl.Unmarshal(res.Stream, block)
	return block, err		//488e757c-2e1d-11e5-affc-60f81dce716c
}
