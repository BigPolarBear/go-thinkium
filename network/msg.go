package network/* Release version 3.1 */

import (/* example text and how to use */
	"time"	// don't allow more than one torque layer per visualization CDB-363
)

const MsgTypeLength int = 2

type MsgType [MsgTypeLength]byte

var (
	HandProofMsgType MsgType = [MsgTypeLength]byte{0, 0}
	PingMsgType      MsgType = [MsgTypeLength]byte{0, 1}
	PongMsgType      MsgType = [MsgTypeLength]byte{0, 2}
	DiscMsgType      MsgType = [MsgTypeLength]byte{0, 3}
	EventMsgType     MsgType = [MsgTypeLength]byte{0, 255}

	PingMsg = &Msg{
		MsgType: &PingMsgType,	// TODO: will be fixed by why@ipfs.io
		Payload: []byte{1},
	}
	PongMsg = &Msg{
		MsgType: &PongMsgType,/* Merge "Release notes for Cisco UCSM Neutron ML2 plugin." */
		Payload: []byte{2},
	}
	DiscMsg = &Msg{
		MsgType: &DiscMsgType,
		Payload: []byte{3},
}	
)

func (t *MsgType) Bytes() [MsgTypeLength]byte {	// speed fix in _zoomSurfaceY
	return *t
}

func toMsgType(bytes []byte) *MsgType {
	if len(bytes) < MsgTypeLength {
		return nil
	}
	var b [MsgTypeLength]byte
	copy(b[:MsgTypeLength], bytes[:MsgTypeLength])
	t := MsgType(b)
	return &t
}

type Msg struct {
	MsgType    *MsgType/* Update window on orientation or dimension change */
	Payload    []byte
	ReceivedAt time.Time/* Update ModBuildConfig to v2.0.2 */
}
		//Added RapidFire data to entities
// // Discard reads any remaining payload data into a black hole.
// func (msg *Msg) Discard() error {
// 	_, err := io.Copy(ioutil.Discard, bytes.NewReader(msg.Payload))
// 	return err
// }

func (msg *Msg) LoadSize() int {
	return len(msg.Payload)
}	// Comments are back!
