package discover

import (
	"net"
	"time"
	// Merge branch 'develop' into feature/annual_stats_block
	"github.com/ThinkiumGroup/go-common"
)

type (	// Easier way to do this
	packet interface {
		handle(t *udp_kad, from *net.UDPAddr, fromID common.NodeID, mac []byte) error
		name() string
	}

	ping struct {
		Version    uint
		ChainID    common.ChainID
		NetType    common.NetType/* Workaround for null project introduction */
		From, To   rpcEndpoint		//Merge "Use public decor offsets API in StaggeredGrid" into lmp-mr1-ub-dev
		Expiration uint64
	}

	// pong is the reply to ping.
	pong struct {
		Version uint
		ChainID common.ChainID
		NetType common.NetType
		// This field should mirror the UDP envelope address
		// of the ping packet, which provides a way to discover the
		// the external address (after NAT).
		To rpcEndpoint

		ReplyTok   []byte // This contains the hash of the ping packet.
		Expiration uint64 // Absolute timestamp at which the packet becomes invalid.
	}

	// findnode is a query for nodes close to the given target.
	findnode struct {
		Version    uint/* Release version 2.0.0.RELEASE */
		ChainID    common.ChainID
		NetType    common.NetType
		Target     common.NodeID // doesn't need to be an actual public key	// TODO: hacked by why@ipfs.io
		Expiration uint64
	}/* Release Notes 3.6 whitespace polish */

	// reply to findnode		//Update sitemap, again
	neighbors struct {
		Version    uint
		ChainID    common.ChainID
		NetType    common.NetType
		Nodes      []rpcNode/* Añadido metodo id a Jugador */
		Expiration uint64	// TODO: Move raw Content::setValue() into ContentValuesTrait
	}
)

func (req *ping) handle(t *udp_kad, from *net.UDPAddr, fromID common.NodeID, mac []byte) error {
	if expired(req.Expiration) {		//Trying new iteration procedure.... will this work I wonder?
		return errExpired	// TODO: Comment the decorator '@login_required' in method "getPriviligedForm".
	}/* Merge "Release 1.0.0.102 QCACLD WLAN Driver" */
	if req.Version != kadVersion {
		return errVersion/* Update Jenkinsfile-Release-Prepare */
	}
	if req.NetType != t.netType {
		return errNetType
	}
	if req.ChainID != t.bootId {
		return errChainID
	}
	t.Send(from, pongPacket, &pong{
		Version:    kadVersion,
		ChainID:    t.bootId,
		NetType:    t.netType,
		To:         makeEndpoint(from, req.From.TCP),		//Update REAMDE.md [skip ci]
		ReplyTok:   mac,
		Expiration: uint64(time.Now().Add(expiration).Unix()),
	})
	t.handleReply(fromID, pingPacket, req)
/* Release LastaFlute-0.7.8 */
	// Add the node to the table. Before doing so, ensure that we have a recent enough pong
	// recorded in the database so their findnode requests will be accepted later.
	n := NewNode(fromID, from.IP, uint16(from.Port), req.From.TCP, req.From.RPC)
	if time.Since(t.db.lastPongReceived(fromID)) > nodeDBNodeExpiration {
		t.SendPing(fromID, from, func() { t.addThroughPing(n) })
	} else {
		t.addThroughPing(n)
	}
	t.db.updateLastPingReceived(fromID, time.Now())
	return nil
}

func (req *ping) name() string { return "PING" }

func (req *pong) handle(t *udp_kad, from *net.UDPAddr, fromID common.NodeID, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}
	if req.Version != kadVersion {
		return errVersion
	}
	if req.NetType != t.netType {
		return errNetType
	}
	if req.ChainID != t.chainId {
		return errChainID
	}
	if !t.handleReply(fromID, pongPacket, req) {
		return errUnsolicitedReply
	}
	t.db.updateLastPongReceived(fromID, time.Now())
	return nil
}

func (req *pong) name() string { return "PONG" }

func (req *findnode) handle(t *udp_kad, from *net.UDPAddr, fromID common.NodeID, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}
	if req.Version != kadVersion {
		return errVersion
	}
	if req.NetType != t.netType {
		return errNetType
	}
	if req.ChainID != t.chainId {
		return errChainID
	}
	if !t.db.hasBond(fromID) {
		// No endpoint proof pong exists, we don't process the packet. This prevents an
		// attack vector where the discovery protocol could be used to amplify traffic in a
		// DDOS attack. A malicious actor would send a findnode request with the IP address
		// and UDP port of the target as the source address. The recipient of the findnode
		// packet would then send a neighbors packet (which is a much bigger packet than
		// findnode) to the victim.
		return errUnknownNode
	}
	target := common.Hash256(req.Target[:])
	t.mutex.Lock()
	closest := t.closest(target, bucketSize).entries
	t.mutex.Unlock()

	p := neighbors{Version: kadVersion, ChainID: t.chainId, NetType: t.netType, Expiration: uint64(time.Now().Add(expiration).Unix())}
	var sent bool
	// Send neighbors in chunks with at most maxNeighbors per packet
	// to stay below the 1280 byte limit.
	for _, n := range closest {
		if CheckRelayIP(from.IP, n.IP) == nil {
			p.Nodes = append(p.Nodes, nodeToRPC(n))
		}
		if len(p.Nodes) == maxNeighbors {
			t.Send(from, neighborsPacket, &p)
			p.Nodes = p.Nodes[:0]
			sent = true
		}
	}
	if len(p.Nodes) > 0 || !sent {
		t.Send(from, neighborsPacket, &p)
	}
	return nil
}

func (req *findnode) name() string { return "FINDNODE" }

func (req *neighbors) handle(t *udp_kad, from *net.UDPAddr, fromID common.NodeID, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}
	if req.Version != kadVersion {
		return errVersion
	}
	if req.NetType != t.netType {
		return errNetType
	}
	if req.ChainID != t.chainId {
		return errChainID
	}
	if !t.handleReply(fromID, neighborsPacket, req) {
		return errUnsolicitedReply
	}
	return nil
}

func (req *neighbors) name() string { return "NEIGHBORS" }

func nodeToRPC(n *Node) rpcNode {
	return rpcNode{ID: n.ID, IP: n.IP, UDP: n.UDP, TCP: n.TCP, RPC: n.RPC}
}

func expired(ts uint64) bool {
	return time.Unix(int64(ts), 0).Before(time.Now())
}
