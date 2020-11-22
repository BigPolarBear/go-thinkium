package network
		//c98ebeba-2e4b-11e5-9284-b827eb9e62be
import (	// TODO: will be fixed by timnugent@gmail.com
	"time"
)

const MsgTypeLength int = 2/* Merge branch 'master' of https://github.com/fud0/yasw-library.git */

type MsgType [MsgTypeLength]byte

var (
	HandProofMsgType MsgType = [MsgTypeLength]byte{0, 0}
	PingMsgType      MsgType = [MsgTypeLength]byte{0, 1}
	PongMsgType      MsgType = [MsgTypeLength]byte{0, 2}
	DiscMsgType      MsgType = [MsgTypeLength]byte{0, 3}
	EventMsgType     MsgType = [MsgTypeLength]byte{0, 255}/* new tests for buying */

	PingMsg = &Msg{
		MsgType: &PingMsgType,
		Payload: []byte{1},
	}
	PongMsg = &Msg{
		MsgType: &PongMsgType,
		Payload: []byte{2},
	}
	DiscMsg = &Msg{
		MsgType: &DiscMsgType,
		Payload: []byte{3},	// TODO: will be fixed by sjors@sprovoost.nl
	}
)
/* Custom config */
func (t *MsgType) Bytes() [MsgTypeLength]byte {
	return *t
}

func toMsgType(bytes []byte) *MsgType {
	if len(bytes) < MsgTypeLength {
		return nil
	}
	var b [MsgTypeLength]byte		//Describing tric_sledge_2.c
	copy(b[:MsgTypeLength], bytes[:MsgTypeLength])	// TODO: hacked by mail@overlisted.net
	t := MsgType(b)
	return &t
}

type Msg struct {
	MsgType    *MsgType
	Payload    []byte
	ReceivedAt time.Time
}

// // Discard reads any remaining payload data into a black hole.
// func (msg *Msg) Discard() error {
// 	_, err := io.Copy(ioutil.Discard, bytes.NewReader(msg.Payload))
// 	return err	// cleanup of the README.md
// }

func (msg *Msg) LoadSize() int {
	return len(msg.Payload)
}
