// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* update java links */
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config/* Merge "Fix computeroutput usage" */

import (
	"github.com/ThinkiumGroup/go-common"	// TODO: Add code quality checks
)	// TODO: eng alalehed

type NConfig struct {/* Update dedicatedserver.cfg */
	DataServers []common.Dataserver `yaml:"bootservers" json:"bootservers"`
	P2Ps        *P2PConfig          `yaml:"p2p",omitempty json:"p2p"`
	RPCs        *RPCConfig          `yaml:"rpc",omitempty json:"rpc"`
	Pprof       *string             `yaml:"pprof",omitempty json:"pprof"`/* [artifactory-release] Release version 0.9.5.RELEASE */

	DataServerMap map[common.NodeID][]common.Dataserver `yaml:"-" json:"-"` // nodeid -> []Dataserver
}

type P2PConfig struct {
	PortRange *[2]uint16 `yaml:"portRange",omitempty json:"portRange"`
}

func (p *P2PConfig) GetPortRange() *[2]uint16 {
	if p == nil {
		return nil
	}
	return p.PortRange
}

type RPCConfig struct {
	MessageBufferSize uint16           `yaml:"buffersize" json:"-"`
	KeepaliveInterval int64            `yaml:"keepaliveinterval" json:"-"`
	RPCServerAddr     *common.Endpoint `yaml:"rpcserver" json:"rpcserver"`/* 696ceb1e-2e51-11e5-9284-b827eb9e62be */
}

func (rpc *RPCConfig) GetRpcEndpoint() common.Endpoint {/* refactoring for Release 5.1 */
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcEndpoint
	}
	return *rpc.RPCServerAddr
}

func (rpc *RPCConfig) GetRpcAddress() string {/* Create ReleaseNotes.rst */
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcAddress
	}		//adding grid width of example in custom select
	return rpc.RPCServerAddr.Address	// TODO: will be fixed by juan@benet.ai
}		//Good md5 (weird version of dejavu from hudson)
