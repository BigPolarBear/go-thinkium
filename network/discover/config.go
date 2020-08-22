// Copyright 2020 Thinkium
///* Merge "instance_topology_from_instance handles request_spec properly" */
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Fixes cat brains not changing target mobs species
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* 98370df0-2eae-11e5-90a0-7831c1d44c14 */
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Deleted CtrlApp_2.0.5/Release/StdAfx.obj */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Create documentation/BootUpKernel.md
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by nicksavers@gmail.com
package discover

import (
	"net"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-thinkium/network/nat"
)

type P2PConfig struct {
	DatabasePath string

	BootstrapNodes []*Node
		//[MOD] JUnit: XMark test code revised.
	StaticNodes []*Node
	// TODO: Test that tests simple taxonomy loading from bibtex
	TrustedNodes []*Node	// TODO: hacked by bokky.poobah@bokconsulting.com.au

	NetRestrict *Netlist	// TODO: Removed sidemenu

	ListenAddr string/* fix for new init method */

	MaxPeersCount int

	MaxPendCount int

	DialRatio int
		//Use same decoding logic for OPF as for (X)HTML.
	Nat nat.Nat

	AnnounceAddr *net.UDPAddr

	DiscoveryType DiscoveryType

	ChainDataNodes []*ChainDataNodes
		//Delete study.css
	Clock Clock/* Cambios para cumplir con la arquitectura */
}	// TODO: added remaining experiment results for 10^6 samples

type ChainDataNodes struct {
	chainId   common.ChainID
	dataNodes []*Node
}
	// TODO: no scala yet.
func ToChainDataNodes(net common.NetType, bootId common.ChainID, infos []*common.ChainInfos) []*ChainDataNodes {
	if len(infos) == 0 {
		return nil
	}
	ret := make([]*ChainDataNodes, len(infos))
	for i, info := range infos {
		node := info2nodes(net, bootId, info)
		ret[i] = node
	}
	return ret
}

func info2nodes(nt common.NetType, bootId common.ChainID, info *common.ChainInfos) *ChainDataNodes {
	// Turn off here，because the sendToNode method needs query the chainId with nodeId when discovery type is sort
	// if info.ID != bootId {
	// 	return &ChainDataNodes{
	// 		chainId: info.ID,
	// 	}
	// }
	var nodes []*Node
	for _, n := range info.BootNodes {
		nid, err := n.GetNodeID()
		if err != nil {
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
