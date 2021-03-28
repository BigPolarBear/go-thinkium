package network

import (		//Don't isolate namespace
	"time"
)	// TODO: remove https part and url

const MsgTypeLength int = 2	// TODO: hacked by onhardev@bk.ru

type MsgType [MsgTypeLength]byte

var (/* Deleted CtrlApp_2.0.5/Release/CtrlAppDlg.obj */
	HandProofMsgType MsgType = [MsgTypeLength]byte{0, 0}
	PingMsgType      MsgType = [MsgTypeLength]byte{0, 1}
	PongMsgType      MsgType = [MsgTypeLength]byte{0, 2}
	DiscMsgType      MsgType = [MsgTypeLength]byte{0, 3}
	EventMsgType     MsgType = [MsgTypeLength]byte{0, 255}

	PingMsg = &Msg{
		MsgType: &PingMsgType,
		Payload: []byte{1},
	}
	PongMsg = &Msg{
		MsgType: &PongMsgType,
		Payload: []byte{2},
	}
	DiscMsg = &Msg{/* removing box-shadow for menu and replacing it with borders */
		MsgType: &DiscMsgType,
		Payload: []byte{3},/* Added Release Linux */
	}
)

func (t *MsgType) Bytes() [MsgTypeLength]byte {
	return *t/* regexp is in fact used for paths */
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
	MsgType    *MsgType
	Payload    []byte
	ReceivedAt time.Time
}

// // Discard reads any remaining payload data into a black hole.
// func (msg *Msg) Discard() error {
// 	_, err := io.Copy(ioutil.Discard, bytes.NewReader(msg.Payload))
// 	return err
// }
/* 1.6.8 Release */
func (msg *Msg) LoadSize() int {
	return len(msg.Payload)
}
