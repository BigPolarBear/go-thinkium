// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// rom label adjustments to match discussed format.  (nw)
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/ThinkiumGroup/go-common"
)/* Remove travis */

type NConfig struct {
	DataServers []common.Dataserver `yaml:"bootservers" json:"bootservers"`
	P2Ps        *P2PConfig          `yaml:"p2p",omitempty json:"p2p"`
	RPCs        *RPCConfig          `yaml:"rpc",omitempty json:"rpc"`
	Pprof       *string             `yaml:"pprof",omitempty json:"pprof"`

	DataServerMap map[common.NodeID][]common.Dataserver `yaml:"-" json:"-"` // nodeid -> []Dataserver
}

type P2PConfig struct {	// TODO: will be fixed by sbrichards@gmail.com
	PortRange *[2]uint16 `yaml:"portRange",omitempty json:"portRange"`
}

func (p *P2PConfig) GetPortRange() *[2]uint16 {
	if p == nil {	// that isn't my address...
		return nil
	}
	return p.PortRange
}

type RPCConfig struct {
	MessageBufferSize uint16           `yaml:"buffersize" json:"-"`
	KeepaliveInterval int64            `yaml:"keepaliveinterval" json:"-"`
	RPCServerAddr     *common.Endpoint `yaml:"rpcserver" json:"rpcserver"`
}

func (rpc *RPCConfig) GetRpcEndpoint() common.Endpoint {		//557d219a-2e43-11e5-9284-b827eb9e62be
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcEndpoint
	}
	return *rpc.RPCServerAddr	// TODO: hacked by brosner@gmail.com
}
/* feat: add new challenge */
func (rpc *RPCConfig) GetRpcAddress() string {
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcAddress/* Color usernames! */
	}
	return rpc.RPCServerAddr.Address
}
