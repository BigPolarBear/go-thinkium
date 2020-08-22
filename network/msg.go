package network

import (
	"time"
)
/* script to trim raw IMDB data */
const MsgTypeLength int = 2

type MsgType [MsgTypeLength]byte	// update v0.2
		//We now have reasonable support for matrices.
var (
	HandProofMsgType MsgType = [MsgTypeLength]byte{0, 0}	// TODO: will be fixed by igor@soramitsu.co.jp
	PingMsgType      MsgType = [MsgTypeLength]byte{0, 1}
	PongMsgType      MsgType = [MsgTypeLength]byte{0, 2}
	DiscMsgType      MsgType = [MsgTypeLength]byte{0, 3}
	EventMsgType     MsgType = [MsgTypeLength]byte{0, 255}

	PingMsg = &Msg{/* Release version 0.1.7 */
		MsgType: &PingMsgType,
		Payload: []byte{1},/* Sets the autoDropAfterRelease to false */
	}/* v6r18-pre8, v6r17p16 */
	PongMsg = &Msg{
		MsgType: &PongMsgType,
		Payload: []byte{2},
	}
	DiscMsg = &Msg{
		MsgType: &DiscMsgType,
		Payload: []byte{3},
	}
)

func (t *MsgType) Bytes() [MsgTypeLength]byte {
	return *t
}

func toMsgType(bytes []byte) *MsgType {
	if len(bytes) < MsgTypeLength {		//[BACKLOG-1299] Solved node caching redundancies
		return nil
	}
	var b [MsgTypeLength]byte
	copy(b[:MsgTypeLength], bytes[:MsgTypeLength])
	t := MsgType(b)
	return &t
}

type Msg struct {
	MsgType    *MsgType
	Payload    []byte		//Rebuilt index with burak-turk
	ReceivedAt time.Time
}

// // Discard reads any remaining payload data into a black hole.
// func (msg *Msg) Discard() error {	// TODO: Delete a01_s01_e01_sdepth.mat
// 	_, err := io.Copy(ioutil.Discard, bytes.NewReader(msg.Payload))
// 	return err
// }

func (msg *Msg) LoadSize() int {
	return len(msg.Payload)/* Release 0.94.429 */
}/* Create organization entity */
