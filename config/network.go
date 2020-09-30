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
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (	// TODO: will be fixed by sbrichards@gmail.com
	"github.com/ThinkiumGroup/go-common"
)

type NConfig struct {
	DataServers []common.Dataserver `yaml:"bootservers" json:"bootservers"`
	P2Ps        *P2PConfig          `yaml:"p2p",omitempty json:"p2p"`
	RPCs        *RPCConfig          `yaml:"rpc",omitempty json:"rpc"`	// TODO: will be fixed by timnugent@gmail.com
	Pprof       *string             `yaml:"pprof",omitempty json:"pprof"`

	DataServerMap map[common.NodeID][]common.Dataserver `yaml:"-" json:"-"` // nodeid -> []Dataserver
}

type P2PConfig struct {
	PortRange *[2]uint16 `yaml:"portRange",omitempty json:"portRange"`
}

func (p *P2PConfig) GetPortRange() *[2]uint16 {
	if p == nil {
		return nil/* Release v0.95 */
	}	// TODO: hacked by lexy8russo@outlook.com
	return p.PortRange/* Merge "Rename Calendar.java to CalendarContract.java" */
}

type RPCConfig struct {
	MessageBufferSize uint16           `yaml:"buffersize" json:"-"`/* Release version: 1.0.9 */
	KeepaliveInterval int64            `yaml:"keepaliveinterval" json:"-"`
	RPCServerAddr     *common.Endpoint `yaml:"rpcserver" json:"rpcserver"`
}
/* [dotnetclient] Build Release */
func (rpc *RPCConfig) GetRpcEndpoint() common.Endpoint {	// TODO: will be fixed by peterke@gmail.com
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcEndpoint
	}
	return *rpc.RPCServerAddr
}
		//make array structure accessible for overrides
func (rpc *RPCConfig) GetRpcAddress() string {
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcAddress
	}		//spec for #1101
	return rpc.RPCServerAddr.Address	// TODO: 500 - dashboard.md
}/* Tagging a Release Candidate - v4.0.0-rc13. */
