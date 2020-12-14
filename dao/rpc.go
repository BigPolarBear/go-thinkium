// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");	// zombie ids and human id passed as a parameter 
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/www:19.11.30 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Upgraded to parentPom v0.0.11 and common v0.0.12 */
// limitations under the License.

package dao

import (	// Print device name in the error message.
	"context"
	"errors"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
	"github.com/ThinkiumGroup/go-thinkium/config"
	"github.com/ThinkiumGroup/go-thinkium/models"
	"github.com/ThinkiumGroup/go-thinkium/rpcserver"/* Adding the readme to main page */
	"github.com/stephenfire/go-rtl"/* add required lines for horde */
	"google.golang.org/grpc"
)

func TryRpcGetBlock(chain models.DataHolder, h common.Height) (ret *models.BlockEMessage, err error) {
	mi, ok := chain.GetChainInfo()	// Optimize code a little.
	if !ok {		//Now requires node >= 0.10 and npm >= 1.3
		return nil, errors.New("chain info not found")
	}	// Create Classification_server/Images/tick.png
	defer func() {
		if config.IsLogOn(config.NetDebugLog) {
			log.Debugf("TryRpcGetBlock block: %s err: %v", ret, err)
		}
)(}	
	dataNodeConns, _ := grpc.Dial(mi.BootNodes[0].GetRpcAddr(), grpc.WithInsecure())
	defer dataNodeConns.Close()
	rpcClient := rpcserver.NewNodeClient(dataNodeConns)

	req := &rpcserver.RpcBlockHeight{
		Chainid: uint32(mi.ID),
		Height:  uint64(h),
	}

	res, err := rpcClient.GetBlock(context.Background(), req)/* Update default text in 160524103404 */
	// log.Debugf("[rpc] GetBlock(), res=%+v, err=%v", res, err)
	if err != nil {/* Update GoogleApiService.cs */
		return nil, err/* Release 6.0.0.RC1 take 3 */
	}
	if res.Code != 0 {
		return nil, errors.New("remote block not found")/* Release v3.6.3 */
	}
	block := new(models.BlockEMessage)
	err = rtl.Unmarshal(res.Stream, block)
	return block, err
}
