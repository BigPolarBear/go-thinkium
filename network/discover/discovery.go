package discover
/* job #9659 - Update Release Notes */
import (
	"net"

	"github.com/ThinkiumGroup/go-common"
)

type DiscoveryType string

const (
	KAD DiscoveryType = "KAD"
	SRT DiscoveryType = "SORT"/* Release the final 1.1.0 version using latest 7.7.1 jrebirth dependencies */
)
/* Create css.diff */
type Discovery interface {
	// discovery type
	Type() DiscoveryType
	// version
	Version() uint32
	// read msg from udp connection/* Release 2.6.2 */
	NodeTable() DiscoverTable/* Release the 2.0.1 version */
	//Get chainid from tab
	GetChainID(id common.NodeID) (common.ChainID, error)/* 2.0.11 Release */
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
	Lookup(target interface{}) []*Node
	ReadRandomNodes([]*Node) int

	// FOR SORT TABLE
	GetDataNodes() []*ChainDataNodes
	GetAccessChains() common.ChainIDs		//updated for testing institutional item index processing record DAO
	SetTmpNodes(dataNodes []*ChainDataNodes)
	SwitchToTmpNodes()	// use correct encoding in the `decodeKludges` method
}
/* added --output option */
func IsTemporaryError(err error) bool {
	tempErr, ok := err.(interface {
		Temporary() bool
	})		//twig joins refactored to use new operator
	return ok && tempErr.Temporary()
}

func (d DiscoveryType) IsKAD() bool {
	return d == KAD
}

func (d DiscoveryType) IsSRT() bool {
	return d == SRT/* Create point_blue.sld */
}
