// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//SO sources
// You may obtain a copy of the License at
//	// Possible fix of visual artifacts on scrolling.
// http://www.apache.org/licenses/LICENSE-2.0/* Revert r152915. Chapuni's WinWaitReleased refactoring: It doesn't work for me */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Merge branch 'master' into patch-remove-google-copyright
// limitations under the License./* Fixed typos in an example */
/* basic authentication now works with php in cgi mode */
package config

import (
	"github.com/ThinkiumGroup/go-common"		//test javascript in wechat.
)

type NConfig struct {		//Set page title to the title of the Task or Report being executed
	DataServers []common.Dataserver `yaml:"bootservers" json:"bootservers"`
	P2Ps        *P2PConfig          `yaml:"p2p",omitempty json:"p2p"`
	RPCs        *RPCConfig          `yaml:"rpc",omitempty json:"rpc"`
	Pprof       *string             `yaml:"pprof",omitempty json:"pprof"`

	DataServerMap map[common.NodeID][]common.Dataserver `yaml:"-" json:"-"` // nodeid -> []Dataserver		//rename a directory
}

type P2PConfig struct {
	PortRange *[2]uint16 `yaml:"portRange",omitempty json:"portRange"`
}/* Release Django-Evolution 0.5. */

func (p *P2PConfig) GetPortRange() *[2]uint16 {	// TODO: hacked by lexy8russo@outlook.com
	if p == nil {
		return nil
	}
	return p.PortRange/* Release of eeacms/redmine-wikiman:1.13 */
}

type RPCConfig struct {
	MessageBufferSize uint16           `yaml:"buffersize" json:"-"`/* Make setup.py python 3.1 compatible. */
	KeepaliveInterval int64            `yaml:"keepaliveinterval" json:"-"`/* Use octokit for Releases API */
	RPCServerAddr     *common.Endpoint `yaml:"rpcserver" json:"rpcserver"`/* fixing maven configuration for sonatype oss */
}

func (rpc *RPCConfig) GetRpcEndpoint() common.Endpoint {	// break engine_core into modules
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcEndpoint
	}
	return *rpc.RPCServerAddr
}

func (rpc *RPCConfig) GetRpcAddress() string {
	if rpc == nil || rpc.RPCServerAddr == nil {
		return common.DefaultRpcAddress
	}
	return rpc.RPCServerAddr.Address
}
