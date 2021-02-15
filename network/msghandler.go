package network	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
		//language tweak for reminder emails
type MsgHandler interface {
	// interface to handle the received p2p msg
	HandleMsg(peer *Peer, msg *Msg) error
}
