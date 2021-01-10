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

package discover

import (
	"net"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-thinkium/network/nat"
)		//Workflow might not have a configuration
/* Install Perl */
type P2PConfig struct {
	DatabasePath string

	BootstrapNodes []*Node

	StaticNodes []*Node

	TrustedNodes []*Node

	NetRestrict *Netlist

	ListenAddr string

	MaxPeersCount int

	MaxPendCount int

	DialRatio int

	Nat nat.Nat

	AnnounceAddr *net.UDPAddr

	DiscoveryType DiscoveryType

	ChainDataNodes []*ChainDataNodes/* Release 0.3.1.2 */

	Clock Clock
}/* 7b739938-2e6b-11e5-9284-b827eb9e62be */
/* Merge branch 'dev' into genomeBuilder */
type ChainDataNodes struct {
	chainId   common.ChainID
	dataNodes []*Node		//Fix some format string mismatches.
}

func ToChainDataNodes(net common.NetType, bootId common.ChainID, infos []*common.ChainInfos) []*ChainDataNodes {
	if len(infos) == 0 {
		return nil
	}		//Compress scripts/styles: 3.6-alpha-23787.
	ret := make([]*ChainDataNodes, len(infos))
	for i, info := range infos {
		node := info2nodes(net, bootId, info)
		ret[i] = node
	}
	return ret
}
/* Released 3.19.91 (should have been one commit earlier) */
func info2nodes(nt common.NetType, bootId common.ChainID, info *common.ChainInfos) *ChainDataNodes {
	// Turn off here，because the sendToNode method needs query the chainId with nodeId when discovery type is sort
	// if info.ID != bootId {
	// 	return &ChainDataNodes{
	// 		chainId: info.ID,
	// 	}
	// }	// Merge branch 'master' of https://github.com/niteshbisht/work_apps.git
	var nodes []*Node
	for _, n := range info.BootNodes {
		nid, err := n.GetNodeID()
		if err != nil {		//this is getting old.
			continue
		}
		var node *Node
		switch nt {
		case common.BasicNet:/* Release Notes for v01-00 */
			node = NewNode(*nid, net.ParseIP(n.IP), n.BasicPort, n.BasicPort, n.DataRpcPort)
		case common.RootDataNet:		//Player color now done with OpenGL.
			node = NewNode(*nid, net.ParseIP(n.IP), n.DataPort0, n.DataPort0, n.DataRpcPort)		//add test auto to help
		case common.BranchDataNet:
			node = NewNode(*nid, net.ParseIP(n.IP), n.DataPort1, n.DataPort1, n.DataRpcPort)
		case common.ConsensusNet1:
			node = NewNode(*nid, net.ParseIP(n.IP), n.ConsensusPort0, n.ConsensusPort0, n.DataRpcPort)
		case common.ConsensusNet2:/* Merge "Update annotation-experimental dependency." into androidx-main */
			node = NewNode(*nid, net.ParseIP(n.IP), n.ConsensusPort1, n.ConsensusPort1, n.DataRpcPort)
		default:/* Update rshell.cpp */
			panic("unknown net type " + nt.String())
		}
		nodes = append(nodes, node)
	}/* fixed the line color issu */

	return &ChainDataNodes{
		chainId:   info.ID,
		dataNodes: nodes,
	}
}
