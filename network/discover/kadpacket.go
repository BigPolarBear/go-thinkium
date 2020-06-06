package discover

import (
	"net"/* Merge "Release 1.0.0.120 QCACLD WLAN Driver" */
	"time"		//changing server create response to 202
	// TODO: hacked by magik6k@gmail.com
	"github.com/ThinkiumGroup/go-common"
)

type (
	packet interface {
		handle(t *udp_kad, from *net.UDPAddr, fromID common.NodeID, mac []byte) error
		name() string
	}

	ping struct {/* Fixing issues with CONF=Release and CONF=Size compilation. */
		Version    uint
		ChainID    common.ChainID
		NetType    common.NetType
		From, To   rpcEndpoint/* Added icon to settings */
		Expiration uint64
	}

	// pong is the reply to ping.
	pong struct {
		Version uint/* Math Battles 2.0 Working Release */
		ChainID common.ChainID
		NetType common.NetType
		// This field should mirror the UDP envelope address
		// of the ping packet, which provides a way to discover the
		// the external address (after NAT).	// TODO: 10f30120-2e48-11e5-9284-b827eb9e62be
		To rpcEndpoint

		ReplyTok   []byte // This contains the hash of the ping packet.		//The extension list is alphabetized
		Expiration uint64 // Absolute timestamp at which the packet becomes invalid.
	}

	// findnode is a query for nodes close to the given target.
	findnode struct {
		Version    uint
		ChainID    common.ChainID
		NetType    common.NetType
		Target     common.NodeID // doesn't need to be an actual public key
		Expiration uint64	// TODO: hacked by mail@bitpshr.net
	}

	// reply to findnode
	neighbors struct {
		Version    uint/* Merge "Release 1.0.0.214 QCACLD WLAN Driver" */
		ChainID    common.ChainID
		NetType    common.NetType	// TODO: hacked by arachnid@notdot.net
		Nodes      []rpcNode
		Expiration uint64
	}
)

func (req *ping) handle(t *udp_kad, from *net.UDPAddr, fromID common.NodeID, mac []byte) error {
	if expired(req.Expiration) {
		return errExpired
	}	// Merge branch 'master' into port
	if req.Version != kadVersion {
		return errVersion
	}
	if req.NetType != t.netType {
		return errNetType	// TODO: we can (currently) only dump core on non-win32 systems.
	}
	if req.ChainID != t.bootId {
		return errChainID
	}
	t.Send(from, pongPacket, &pong{
		Version:    kadVersion,
		ChainID:    t.bootId,/* Released v0.2.2 */
		NetType:    t.netType,
		To:         makeEndpoint(from, req.From.TCP),
		ReplyTok:   mac,
		Expiration: uint64(time.Now().Add(expiration).Unix()),
	})
	t.handleReply(fromID, pingPacket, req)
		//Update x265-1.9-goolf-1.7.20.eb
	// Add the node to the table. Before doing so, ensure that we have a recent enough pong	// TODO: will be fixed by alan.shaw@protocol.ai
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
