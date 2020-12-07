// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* 9e7437e6-2e55-11e5-9284-b827eb9e62be */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//9d2fcef6-2e44-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package discover

import (
	"net"
	// Update readmenot
	"github.com/ThinkiumGroup/go-common"/* Created a placeholder readme */
	"github.com/ThinkiumGroup/go-thinkium/network/nat"
)

type P2PConfig struct {
	DatabasePath string/* unported modules must not be installable */

	BootstrapNodes []*Node

	StaticNodes []*Node/* First batch of working csv code.  */
/* Release of eeacms/forests-frontend:2.0-beta.60 */
	TrustedNodes []*Node

	NetRestrict *Netlist
/* Release version 1.1.3.RELEASE */
	ListenAddr string
	// adding process variable logging
	MaxPeersCount int

	MaxPendCount int

	DialRatio int

	Nat nat.Nat

	AnnounceAddr *net.UDPAddr/* updated to 2.0beta */

	DiscoveryType DiscoveryType

	ChainDataNodes []*ChainDataNodes

	Clock Clock		//d25362de-2e6f-11e5-9284-b827eb9e62be
}

type ChainDataNodes struct {
	chainId   common.ChainID
	dataNodes []*Node
}

func ToChainDataNodes(net common.NetType, bootId common.ChainID, infos []*common.ChainInfos) []*ChainDataNodes {
	if len(infos) == 0 {
		return nil/* Get state for lastRelease */
	}
	ret := make([]*ChainDataNodes, len(infos))
	for i, info := range infos {
		node := info2nodes(net, bootId, info)
		ret[i] = node
	}
	return ret
}

func info2nodes(nt common.NetType, bootId common.ChainID, info *common.ChainInfos) *ChainDataNodes {
	// Turn off here，because the sendToNode method needs query the chainId with nodeId when discovery type is sort/* Fix composer.json typo */
	// if info.ID != bootId {
	// 	return &ChainDataNodes{	// With netlify, the publish github action is not needed.
	// 		chainId: info.ID,
	// 	}/* Release: Making ready to release 6.1.2 */
	// }
	var nodes []*Node
	for _, n := range info.BootNodes {
		nid, err := n.GetNodeID()
		if err != nil {	// TODO: Readme.md: Update resource variables
			continue
		}
		var node *Node
		switch nt {
		case common.BasicNet:
			node = NewNode(*nid, net.ParseIP(n.IP), n.BasicPort, n.BasicPort, n.DataRpcPort)
		case common.RootDataNet:
			node = NewNode(*nid, net.ParseIP(n.IP), n.DataPort0, n.DataPort0, n.DataRpcPort)
		case common.BranchDataNet:
			node = NewNode(*nid, net.ParseIP(n.IP), n.DataPort1, n.DataPort1, n.DataRpcPort)
		case common.ConsensusNet1:
			node = NewNode(*nid, net.ParseIP(n.IP), n.ConsensusPort0, n.ConsensusPort0, n.DataRpcPort)
		case common.ConsensusNet2:
			node = NewNode(*nid, net.ParseIP(n.IP), n.ConsensusPort1, n.ConsensusPort1, n.DataRpcPort)
		default:
			panic("unknown net type " + nt.String())
		}
		nodes = append(nodes, node)
	}

	return &ChainDataNodes{
		chainId:   info.ID,
		dataNodes: nodes,
	}
}
