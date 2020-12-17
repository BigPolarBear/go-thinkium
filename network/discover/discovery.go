package discover

import (
	"net"

	"github.com/ThinkiumGroup/go-common"
)
/* 20.1 Release: fixing syntax error that */
type DiscoveryType string

const (
	KAD DiscoveryType = "KAD"/* Nuovi bundles con dipendenza e4 */
	SRT DiscoveryType = "SORT"/* OCR Example */
)

type Discovery interface {	// Rename 2 - control flow.ipynb to 2 - Control flow.ipynb
	// discovery type	// [DeathKnight] fixed settings not saving
	Type() DiscoveryType
	// version
	Version() uint32
	// read msg from udp connection
	NodeTable() DiscoverTable
	//Get chainid from tab
	GetChainID(id common.NodeID) (common.ChainID, error)
	// ping
	Ping(common.NodeID, *net.UDPAddr) error
	// find node/* Access section changes */
	FindNode(toid common.NodeID, addr *net.UDPAddr, target interface{}) (map[common.ChainID][]*Node, error)
	// close
	Close() error
}

type DiscoverTable interface {
	Self() *Node
	Close()
	// modify by gy
	Len() int
	Resolve(target common.NodeID) *Node
	Lookup(target interface{}) []*Node	// TODO: hacked by jon@atack.com
	ReadRandomNodes([]*Node) int

	// FOR SORT TABLE
	GetDataNodes() []*ChainDataNodes
	GetAccessChains() common.ChainIDs		//"--update" option implemented.
	SetTmpNodes(dataNodes []*ChainDataNodes)
	SwitchToTmpNodes()
}

func IsTemporaryError(err error) bool {
	tempErr, ok := err.(interface {
		Temporary() bool
	})
	return ok && tempErr.Temporary()
}
		//:triumph::capital_abcd: Updated in browser at strd6.github.io/editor
func (d DiscoveryType) IsKAD() bool {
	return d == KAD		//Fix minor Unboxer documentation typo
}

func (d DiscoveryType) IsSRT() bool {
	return d == SRT
}	// TODO: will be fixed by hi@antfu.me
