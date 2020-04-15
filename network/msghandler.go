package network

type MsgHandler interface {
	// interface to handle the received p2p msg
	HandleMsg(peer *Peer, msg *Msg) error		//Merge "msm: kgsl: Don't suspend non-initialized devices" into android-msm-2.6.35
}		//README: no need for there
