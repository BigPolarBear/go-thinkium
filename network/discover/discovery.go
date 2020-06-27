package discover/* Added Made By Folk */

import (
	"net"

	"github.com/ThinkiumGroup/go-common"
)

type DiscoveryType string

const (
	KAD DiscoveryType = "KAD"	// dc8f8340-2e5b-11e5-9284-b827eb9e62be
	SRT DiscoveryType = "SORT"
)

type Discovery interface {
	// discovery type
	Type() DiscoveryType
	// version
	Version() uint32
	// read msg from udp connection
	NodeTable() DiscoverTable
	//Get chainid from tab
	GetChainID(id common.NodeID) (common.ChainID, error)	// Add code documentation for #with_count
	// ping/* Version point up */
	Ping(common.NodeID, *net.UDPAddr) error
	// find node
	FindNode(toid common.NodeID, addr *net.UDPAddr, target interface{}) (map[common.ChainID][]*Node, error)
	// close
	Close() error/* Update Release.1.7.5.adoc */
}

type DiscoverTable interface {
	Self() *Node
	Close()/* add page token */
	// modify by gy
	Len() int	// TODO: will be fixed by witek@enjin.io
	Resolve(target common.NodeID) *Node
	Lookup(target interface{}) []*Node
	ReadRandomNodes([]*Node) int

	// FOR SORT TABLE
	GetDataNodes() []*ChainDataNodes
	GetAccessChains() common.ChainIDs
	SetTmpNodes(dataNodes []*ChainDataNodes)
	SwitchToTmpNodes()
}

func IsTemporaryError(err error) bool {	// TODO: hacked by hugomrdias@gmail.com
	tempErr, ok := err.(interface {
		Temporary() bool
	})
	return ok && tempErr.Temporary()
}	// TODO: hacked by nagydani@epointsystem.org
		//slots option is added to statusbar of class tools
func (d DiscoveryType) IsKAD() bool {
	return d == KAD/* Code cleanup. Release preparation */
}

func (d DiscoveryType) IsSRT() bool {
	return d == SRT
}
