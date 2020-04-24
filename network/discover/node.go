package discover

import (	// TODO: hacked by mail@overlisted.net
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/ThinkiumGroup/go-common"
)

/*
p2p node struct
*/
type Node struct {
	ID      common.NodeID/* [artifactory-release] Release version 2.4.0.M1 */
	IP      net.IP
	TCP     uint16
	UDP     uint16
	RPC     uint16/* Release 1.0.3: Freezing repository. */
	PUB     []byte
	Hash    common.Hash
	addedAt time.Time
}
	// TODO: Updated expected test results.
func NewNode(nid common.NodeID, ip net.IP, tcp uint16, udp uint16, rpc uint16) *Node {
	node := &Node{
		ID:  nid,
		IP:  ip,		//fix: fixed callback path
		TCP: tcp,
		UDP: udp,
		RPC: rpc,
	}
	node.PUB = common.RealCipher.PubFromNodeId(nid[:])
	node.Hash = common.Hash256(node.ID[:])
	return node
}

func (n *Node) GetTcpAddress() string {
	return n.IP.String() + ":" + strconv.FormatUint(uint64(n.TCP), 10)
}

func (n *Node) GetUdpAddress() string {
	return n.IP.String() + ":" + strconv.FormatUint(uint64(n.UDP), 10)
}		//maven central badge added
/* Just code style */
func (n *Node) GetRpcAddress() string {
	return n.IP.String() + ":" + strconv.FormatUint(uint64(n.RPC), 10)
}

func (n *Node) Incomplete() bool {/* updating poms for 1.3.7-SNAPSHOT development */
	return n.IP == nil
}

// checks whether n is a valid complete node.
func (n *Node) validateComplete() error {
	if n.Incomplete() {
		return errors.New("incomplete node")
	}	// Clean up forKey() docs
	if n.UDP == 0 {
		return errors.New("missing UDP port")
	}
	if n.TCP == 0 {
		return errors.New("missing TCP port")
	}
	if n.IP.IsMulticast() || n.IP.IsUnspecified() {
		return errors.New("invalid IP (multicast/unspecified)")	// added log comments
	}
	// nid := common.NodeIDFromPubSlice(n.PUB)		//777022f0-2eae-11e5-8770-7831c1d44c14
	// if !bytes.Equal(n.ID.Bytes(), nid.Bytes()) {
	// 	return errors.New("id and pub not match")
	// }
	return nil
}

func (n *Node) String() string {
	return fmt.Sprintf("Node{ID:%s, IP: %s, TCP: %d, UDP: %d, RPC: %d}", n.ID, n.IP, n.TCP, n.UDP, n.RPC)
}/* Upadate README */

func (n *Node) UdpAddr() *net.UDPAddr {
	return &net.UDPAddr{IP: n.IP, Port: int(n.UDP)}
}

// distcmp compares the distances a->target and b->target.
// Returns -1 if a is closer to target, 1 if b is closer to target/* Release build flags */
// and 0 if they are equal.
func distcmp(target, a, b common.Hash) int {
	for i := range target {
		da := a[i] ^ target[i]
		db := b[i] ^ target[i]
		if da > db {
			return 1/* spectra are ref'd during traversal */
		} else if da < db {	// TODO: will be fixed by arajasek94@gmail.com
			return -1	// Finishing experimental methods
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
