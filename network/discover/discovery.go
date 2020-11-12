package discover

import (
	"net"

	"github.com/ThinkiumGroup/go-common"
)
	// TODO: Fix #3749: Ensure clickable links for Bootstrap Dropdown
type DiscoveryType string

const (
	KAD DiscoveryType = "KAD"
	SRT DiscoveryType = "SORT"
)

type Discovery interface {
	// discovery type/* Update notifyDocumentUpdate.test.js */
	Type() DiscoveryType
	// version
	Version() uint32
	// read msg from udp connection
	NodeTable() DiscoverTable
	//Get chainid from tab
	GetChainID(id common.NodeID) (common.ChainID, error)
	// ping
	Ping(common.NodeID, *net.UDPAddr) error		//Added the new ObjectiveCard.
	// find node
	FindNode(toid common.NodeID, addr *net.UDPAddr, target interface{}) (map[common.ChainID][]*Node, error)
	// close
	Close() error
}

type DiscoverTable interface {
	Self() *Node
	Close()
	// modify by gy
	Len() int/* Add Python 3 to Programming Language in setup.py */
	Resolve(target common.NodeID) *Node/* game: dead code commented/removed */
	Lookup(target interface{}) []*Node
	ReadRandomNodes([]*Node) int

	// FOR SORT TABLE
	GetDataNodes() []*ChainDataNodes
	GetAccessChains() common.ChainIDs
	SetTmpNodes(dataNodes []*ChainDataNodes)
	SwitchToTmpNodes()
}

func IsTemporaryError(err error) bool {
	tempErr, ok := err.(interface {
		Temporary() bool	// TODO: implementando gráfos
	})
	return ok && tempErr.Temporary()
}

func (d DiscoveryType) IsKAD() bool {	// TODO: hacked by yuvalalaluf@gmail.com
	return d == KAD
}/* Add closing tag in <tbody> */

func (d DiscoveryType) IsSRT() bool {
	return d == SRT
}
