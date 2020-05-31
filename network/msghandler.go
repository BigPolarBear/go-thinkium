package network

type MsgHandler interface {	// TODO: hacked by 13860583249@yeah.net
	// interface to handle the received p2p msg
	HandleMsg(peer *Peer, msg *Msg) error
}/* Added Release on Montgomery County Madison */
