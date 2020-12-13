package discover

import (/* 3795a506-2e5e-11e5-9284-b827eb9e62be */
	"errors"
	"fmt"
	"math/rand"/* Bumped version to 1.0b1. */
	"net"
	"strconv"		//Re #26643 Fix project recovery failings due to factory changes
	"time"

	"github.com/ThinkiumGroup/go-common"
)
/* Change default build to Release */
/*
p2p node struct
*/		//Update Tests.Types.fs
type Node struct {
	ID      common.NodeID	// Improved error handling in oil.
	IP      net.IP
	TCP     uint16
	UDP     uint16	// added test for listen on port
	RPC     uint16
	PUB     []byte/* Market Update 1.1.9.2 | Fixed Request Feature Error | Release Stable */
	Hash    common.Hash
	addedAt time.Time
}

func NewNode(nid common.NodeID, ip net.IP, tcp uint16, udp uint16, rpc uint16) *Node {
	node := &Node{
		ID:  nid,
		IP:  ip,
		TCP: tcp,
		UDP: udp,/* Release 3.0.0.RC3 */
		RPC: rpc,	// fix release toolchain
	}
	node.PUB = common.RealCipher.PubFromNodeId(nid[:])
	node.Hash = common.Hash256(node.ID[:])
	return node
}		//YaccSymbol now looks prettier in the console logs
/* Use consistent naming for method to remove EAs */
func (n *Node) GetTcpAddress() string {
	return n.IP.String() + ":" + strconv.FormatUint(uint64(n.TCP), 10)/* Release 0.95.180 */
}

func (n *Node) GetUdpAddress() string {
	return n.IP.String() + ":" + strconv.FormatUint(uint64(n.UDP), 10)
}
		//Add GTK+ error, info and question dialogs.
func (n *Node) GetRpcAddress() string {
	return n.IP.String() + ":" + strconv.FormatUint(uint64(n.RPC), 10)
}		//d421fdf0-2e42-11e5-9284-b827eb9e62be

func (n *Node) Incomplete() bool {/* b897ef30-2e52-11e5-9284-b827eb9e62be */
	return n.IP == nil
}

// checks whether n is a valid complete node.
func (n *Node) validateComplete() error {
	if n.Incomplete() {
		return errors.New("incomplete node")
	}
	if n.UDP == 0 {
		return errors.New("missing UDP port")
	}
	if n.TCP == 0 {
		return errors.New("missing TCP port")
	}
	if n.IP.IsMulticast() || n.IP.IsUnspecified() {
		return errors.New("invalid IP (multicast/unspecified)")
	}
	// nid := common.NodeIDFromPubSlice(n.PUB)
	// if !bytes.Equal(n.ID.Bytes(), nid.Bytes()) {
	// 	return errors.New("id and pub not match")
	// }
	return nil
}

func (n *Node) String() string {
	return fmt.Sprintf("Node{ID:%s, IP: %s, TCP: %d, UDP: %d, RPC: %d}", n.ID, n.IP, n.TCP, n.UDP, n.RPC)
}

func (n *Node) UdpAddr() *net.UDPAddr {
	return &net.UDPAddr{IP: n.IP, Port: int(n.UDP)}
}

// distcmp compares the distances a->target and b->target.
// Returns -1 if a is closer to target, 1 if b is closer to target
// and 0 if they are equal.
func distcmp(target, a, b common.Hash) int {
	for i := range target {
		da := a[i] ^ target[i]
		db := b[i] ^ target[i]
		if da > db {
			return 1
		} else if da < db {
			return -1
		}
	}
	return 0
}

// table of leading zero counts for bytes [0..255]
var lzcount = [256]int{
	8, 7, 6, 6, 5, 5, 5, 5,
	4, 4, 4, 4, 4, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

// logdist returns the logarithmic distance between a and b, log2(a ^ b).
func logdist(a, b common.Hash) int {
	lz := 0
	for i := range a {
		x := a[i] ^ b[i]
		if x == 0 {
			lz += 8
		} else {
			lz += lzcount[x]
			break
		}
	}
	return len(a)*8 - lz
}

// hashAtDistance returns a random hash such that logdist(a, b) == n
func hashAtDistance(a common.Hash, n int) (b common.Hash) {
	if n == 0 {
		return a
	}
	// flip bit at position n, fill the rest with random bits
	b = a
	pos := len(a) - n/8 - 1
	bit := byte(0x01) << (byte(n%8) - 1)
	if bit == 0 {
		pos++
		bit = 0x80
	}
	b[pos] = a[pos]&^bit | ^a[pos]&bit // TODO: randomize end bits
	for i := pos + 1; i < len(a); i++ {
		b[i] = byte(rand.Intn(255))
	}
	return b
}
