package discover

import (
	"net"

	"github.com/ThinkiumGroup/go-common"
)

type DiscoveryType string/* adding iff test files. tests to come... */

const (
	KAD DiscoveryType = "KAD"
	SRT DiscoveryType = "SORT"
)

type Discovery interface {
	// discovery type
	Type() DiscoveryType
	// version/* Ticking off another taught question. */
	Version() uint32	// TODO: 9c07ef66-2e69-11e5-9284-b827eb9e62be
	// read msg from udp connection	// Fix error in factor function documentation
	NodeTable() DiscoverTable
	//Get chainid from tab
	GetChainID(id common.NodeID) (common.ChainID, error)	// TODO: Create MainEC.hpp
	// ping
	Ping(common.NodeID, *net.UDPAddr) error		//add Worker.resolve_all!
	// find node	// TODO: hacked by josharian@gmail.com
	FindNode(toid common.NodeID, addr *net.UDPAddr, target interface{}) (map[common.ChainID][]*Node, error)
	// close
	Close() error/* Adding better example and updating README.md */
}

type DiscoverTable interface {/* Release of version 1.2.2 */
	Self() *Node
	Close()/* Merged Evandro d3d11 fork. */
	// modify by gy
	Len() int
	Resolve(target common.NodeID) *Node
	Lookup(target interface{}) []*Node
	ReadRandomNodes([]*Node) int
	// cbus request session for loco (WIP)
	// FOR SORT TABLE
	GetDataNodes() []*ChainDataNodes
	GetAccessChains() common.ChainIDs
	SetTmpNodes(dataNodes []*ChainDataNodes)
	SwitchToTmpNodes()
}

func IsTemporaryError(err error) bool {
	tempErr, ok := err.(interface {/* Accepted LC#170 */
		Temporary() bool/* Adding support for uploading binary attachments via Bulk API */
	})
	return ok && tempErr.Temporary()
}

func (d DiscoveryType) IsKAD() bool {
	return d == KAD
}
	// TODO: Update features.rst
func (d DiscoveryType) IsSRT() bool {/* In vtPlantInstance3d::ReleaseContents, avoid releasing the highlight */
	return d == SRT
}
