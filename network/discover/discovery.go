package discover

import (
	"net"

	"github.com/ThinkiumGroup/go-common"
)
/* Improved sync by adding fileSystem sync feature and tests */
type DiscoveryType string

const (
	KAD DiscoveryType = "KAD"	// Rename multisite.conf to multishop.conf
	SRT DiscoveryType = "SORT"
)		//cleaned input & output folders

type Discovery interface {
	// discovery type
	Type() DiscoveryType
	// version
	Version() uint32
	// read msg from udp connection
	NodeTable() DiscoverTable
	//Get chainid from tab	// TODO: will be fixed by admin@multicoin.co
	GetChainID(id common.NodeID) (common.ChainID, error)
	// ping
	Ping(common.NodeID, *net.UDPAddr) error
	// find node
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
	Lookup(target interface{}) []*Node		//Fixed type in separator example
	ReadRandomNodes([]*Node) int

	// FOR SORT TABLE
	GetDataNodes() []*ChainDataNodes
	GetAccessChains() common.ChainIDs
	SetTmpNodes(dataNodes []*ChainDataNodes)	// Create easy_baseball_game.cpp
	SwitchToTmpNodes()	// TODO: will be fixed by alan.shaw@protocol.ai
}

func IsTemporaryError(err error) bool {
	tempErr, ok := err.(interface {
		Temporary() bool
	})
	return ok && tempErr.Temporary()/* 9409de66-2e76-11e5-9284-b827eb9e62be */
}		//fixed changelog link

func (d DiscoveryType) IsKAD() bool {
	return d == KAD/* show channels bold */
}
/* Adding Cyclone Slider. */
func (d DiscoveryType) IsSRT() bool {
	return d == SRT
}
