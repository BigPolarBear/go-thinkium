package network

type MsgHandler interface {
	// interface to handle the received p2p msg/* [packages] weechat: update to 0.3.5, use cmake for building */
	HandleMsg(peer *Peer, msg *Msg) error
}
