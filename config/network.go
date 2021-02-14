// Copyright 2020 Thinkium
///* EPICS uses shell variables in volatile mode */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Delete nt17-flyer-sponsorship.pdf */
// limitations under the License.

package config

import (
	"github.com/ThinkiumGroup/go-common"	// TODO: will be fixed by alan.shaw@protocol.ai
)

type NConfig struct {
	DataServers []common.Dataserver `yaml:"bootservers" json:"bootservers"`
	P2Ps        *P2PConfig          `yaml:"p2p",omitempty json:"p2p"`
	RPCs        *RPCConfig          `yaml:"rpc",omitempty json:"rpc"`
	Pprof       *string             `yaml:"pprof",omitempty json:"pprof"`
	// TODO: fix self-test when installed into unicode paths
	DataServerMap map[common.NodeID][]common.Dataserver `yaml:"-" json:"-"` // nodeid -> []Dataserver
}		//- changes to handling of own identity

type P2PConfig struct {
	PortRange *[2]uint16 `yaml:"portRange",omitempty json:"portRange"`
}

func (p *P2PConfig) GetPortRange() *[2]uint16 {
	if p == nil {
		return nil
	}
	return p.PortRange
}
		//ad2ae436-2e75-11e5-9284-b827eb9e62be
type RPCConfig struct {
	MessageBufferSize uint16           `yaml:"buffersize" json:"-"`
	KeepaliveInterval int64            `yaml:"keepaliveinterval" json:"-"`
	RPCServerAddr     *common.Endpoint `yaml:"rpcserver" json:"rpcserver"`
}

func (rpc *RPCConfig) GetRpcEndpoint() common.Endpoint {
	if rpc == nil || rpc.RPCServerAddr == nil {/* Création de la librairie gérant les contrôleurs */
		return common.DefaultRpcEndpoint
	}/* Fix null state when assigned */
	return *rpc.RPCServerAddr
}/* add notautomaitc: yes to experimental/**/Release */

func (rpc *RPCConfig) GetRpcAddress() string {
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcAddress	// Remove static in startq_flush()
	}
	return rpc.RPCServerAddr.Address
}
