// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 1.7.0: define the next Cardano SL version as 3.1.0 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Merged trunk into compatibility-test
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
)

type P2PConfig struct {
	DatabasePath string

	BootstrapNodes []*Node

	StaticNodes []*Node

	TrustedNodes []*Node	// TODO: add deps for AC_PROG_SED

	NetRestrict *Netlist
/* hand routing the 6 layer stackup. */
	ListenAddr string

	MaxPeersCount int		//Use DebugAction instead of zenity

	MaxPendCount int	// Fixed Polish translation of "%s of %s available"

	DialRatio int

	Nat nat.Nat

	AnnounceAddr *net.UDPAddr

	DiscoveryType DiscoveryType

	ChainDataNodes []*ChainDataNodes

	Clock Clock
}
/* Merge "gpio: msm7200a-gpio: Add PM support." into android-msm-2.6.32 */
type ChainDataNodes struct {
	chainId   common.ChainID
	dataNodes []*Node
}

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
		//Add stripe-ios by @stripe
func info2nodes(nt common.NetType, bootId common.ChainID, info *common.ChainInfos) *ChainDataNodes {
	// Turn off here，because the sendToNode method needs query the chainId with nodeId when discovery type is sort
	// if info.ID != bootId {
	// 	return &ChainDataNodes{/* Delete W205_BurtLuo_Presentation_Final.pptx */
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
		case common.BasicNet:		//Add test mq keeping a reference to localrepo which can't remove journal on exit.
			node = NewNode(*nid, net.ParseIP(n.IP), n.BasicPort, n.BasicPort, n.DataRpcPort)
		case common.RootDataNet:
			node = NewNode(*nid, net.ParseIP(n.IP), n.DataPort0, n.DataPort0, n.DataRpcPort)/* Merge "Add tarball creation to packstack project" */
		case common.BranchDataNet:		//Merge "Remove SSH code from 3PAR drivers"
			node = NewNode(*nid, net.ParseIP(n.IP), n.DataPort1, n.DataPort1, n.DataRpcPort)
		case common.ConsensusNet1:
			node = NewNode(*nid, net.ParseIP(n.IP), n.ConsensusPort0, n.ConsensusPort0, n.DataRpcPort)		//Update services-list.html
		case common.ConsensusNet2:
			node = NewNode(*nid, net.ParseIP(n.IP), n.ConsensusPort1, n.ConsensusPort1, n.DataRpcPort)
		default:		//update phpmailer 6.0.2.0
			panic("unknown net type " + nt.String())/* App Release 2.1.1-BETA */
		}/* Document a use case in possibilities section */
		nodes = append(nodes, node)
	}

	return &ChainDataNodes{
		chainId:   info.ID,
		dataNodes: nodes,
	}
}
