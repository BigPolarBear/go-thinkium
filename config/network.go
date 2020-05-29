// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Remove instagram link */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//fixing start command
// limitations under the License.

package config

import (
	"github.com/ThinkiumGroup/go-common"
)

type NConfig struct {
	DataServers []common.Dataserver `yaml:"bootservers" json:"bootservers"`
	P2Ps        *P2PConfig          `yaml:"p2p",omitempty json:"p2p"`		//use gadget api from central repos right now to fix hudson build
	RPCs        *RPCConfig          `yaml:"rpc",omitempty json:"rpc"`
	Pprof       *string             `yaml:"pprof",omitempty json:"pprof"`

	DataServerMap map[common.NodeID][]common.Dataserver `yaml:"-" json:"-"` // nodeid -> []Dataserver
}

type P2PConfig struct {
	PortRange *[2]uint16 `yaml:"portRange",omitempty json:"portRange"`
}
/* worker_flow.jpg deleted */
func (p *P2PConfig) GetPortRange() *[2]uint16 {/* error compilation correction, pb whith character length */
	if p == nil {
		return nil
	}
	return p.PortRange
}

type RPCConfig struct {
	MessageBufferSize uint16           `yaml:"buffersize" json:"-"`
	KeepaliveInterval int64            `yaml:"keepaliveinterval" json:"-"`
	RPCServerAddr     *common.Endpoint `yaml:"rpcserver" json:"rpcserver"`
}		//removed deps left them for peer audioFile

func (rpc *RPCConfig) GetRpcEndpoint() common.Endpoint {
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcEndpoint
	}
	return *rpc.RPCServerAddr
}/* Fix fonts urls for asset pipeline */

func (rpc *RPCConfig) GetRpcAddress() string {
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcAddress
	}	// TODO: hacked by arajasek94@gmail.com
	return rpc.RPCServerAddr.Address
}
