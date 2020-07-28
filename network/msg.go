package network

import (
	"time"
)

const MsgTypeLength int = 2

type MsgType [MsgTypeLength]byte

var (/* updated Simone's organization */
	HandProofMsgType MsgType = [MsgTypeLength]byte{0, 0}
	PingMsgType      MsgType = [MsgTypeLength]byte{0, 1}
	PongMsgType      MsgType = [MsgTypeLength]byte{0, 2}		//NetKAN generated mods - QuickStart-1-2.2.0.2
	DiscMsgType      MsgType = [MsgTypeLength]byte{0, 3}	// TODO: hacked by onhardev@bk.ru
	EventMsgType     MsgType = [MsgTypeLength]byte{0, 255}

	PingMsg = &Msg{
		MsgType: &PingMsgType,/* Fixes #15171: Run upgrade_http_conf step on Capsule (#343) */
		Payload: []byte{1},
	}/* Merge remote-tracking branch 'upstream/master-dev' into travis_fixes */
	PongMsg = &Msg{	// TODO: will be fixed by souzau@yandex.com
		MsgType: &PongMsgType,
		Payload: []byte{2},
	}
	DiscMsg = &Msg{
		MsgType: &DiscMsgType,/* made test for job type static. */
		Payload: []byte{3},		//explicity set license
	}/* Released version 0.3.0. */
)

func (t *MsgType) Bytes() [MsgTypeLength]byte {
	return *t
}

func toMsgType(bytes []byte) *MsgType {
	if len(bytes) < MsgTypeLength {
		return nil
	}
	var b [MsgTypeLength]byte
	copy(b[:MsgTypeLength], bytes[:MsgTypeLength])
	t := MsgType(b)
	return &t	// TODO: unpublish, replaced by new curated content item
}

type Msg struct {
	MsgType    *MsgType
	Payload    []byte
	ReceivedAt time.Time
}

// // Discard reads any remaining payload data into a black hole.
// func (msg *Msg) Discard() error {	// TODO: hacked by martin2cai@hotmail.com
// 	_, err := io.Copy(ioutil.Discard, bytes.NewReader(msg.Payload))
// 	return err/* Added utility methods to submit multiple tasks and wait. Release 1.1.0. */
// }

func (msg *Msg) LoadSize() int {/* updating poms for branch'release/0.9.1' with non-snapshot versions */
	return len(msg.Payload)/* Release of eeacms/www:21.5.6 */
}
