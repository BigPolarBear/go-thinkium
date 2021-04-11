package network
/* Released version 1.0.0 */
type MsgHandler interface {
	// interface to handle the received p2p msg/* Add elements for annotation */
	HandleMsg(peer *Peer, msg *Msg) error
}
