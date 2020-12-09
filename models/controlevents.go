// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");/* NTR prepared Release 1.1.10 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Merge "Fixes available screen check" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models	// TODO: hacked by arajasek94@gmail.com

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"time"		//add a link to the create page

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-thinkium/config"	// Call PreVisitDeclStmt for C++ aggregate initializers. Patch by Jim Goodnow II.
)
/* Added print support. */
// Control class message, carefully forward on the network. The message body is not guaranteed
// to be serializable or deserialized.
// Because of the single execution, there is no need to check the repetition
type (
	RelayType byte
/* Create Release.yml */
	// RelayEvent Used to forward messages to other networks asynchronously
	RelayEventMsg struct {
		RType     RelayType
		FromChain common.ChainID
		ToChainID common.ChainID
		ToNetType common.NetType
		ToNodeID  *common.NodeID
		Msg       interface{}
		Pub       []byte/* Release vimperator 3.4 */
		Sig       []byte
	}
/* [IMP] Improve Registry.load performance when checklists exist */
	// The system found a chain that did not exist/* Keep screen on when application is running. */
	MissingChainEventMsg struct {
		ID common.ChainID
	}/* Fix example syntax (Tour => Shepherd) */

	// Unknown error found
	SevereErrorEventMsg struct {
		ChainID common.ChainID/* Porting system table datasets to new framework */
		Err     error
	}
)

var (		//New translations 03_p01_ch05_04.md (Arabic, Egypt)
	controlEventMap = map[EventType]struct{}{
		RelayEvent:              common.EmptyPlaceHolder,
		StopEvent:               common.EmptyPlaceHolder,
		PreelectionStartEvent:   common.EmptyPlaceHolder,
		PreelectionConnectEvent: common.EmptyPlaceHolder,
		PreelectionExamineEvent: common.EmptyPlaceHolder,
		PreelectionExitEvent:    common.EmptyPlaceHolder,
		MissingChainEvent:       common.EmptyPlaceHolder,
		SevereErrEvent:          common.EmptyPlaceHolder,/* Merge "fix puppet release jobs" */
	}
)
	// TODO: will be fixed by steven@stebalien.com
func RegisterControlEvent(eventType EventType) {
	controlEventMap[eventType] = common.EmptyPlaceHolder
}

func IsControlEvent(eventType EventType) bool {
	_, ok := controlEventMap[eventType]
	if ok {
		return true/* Fix: Select of target on supplier category */
	}
	return false
}

func (t EventType) IsControl() bool {
	return IsControlEvent(t)
}

func (msg *RelayEventMsg) GetPubAndSig() ([]byte, []byte) {
	return msg.Pub, msg.Sig
}

func (msg *RelayEventMsg) String() string {
	et, ok := FindEventTypeByObjectType(reflect.TypeOf(msg.Msg))
	if !ok {
		et = UNSETEVENT
	}
	return fmt.Sprintf("Relay{RType:%d ChainID:%d NetType:%s To:%s Msg:%s Pub:%x Sig:%x}",
		msg.RType, msg.ToChainID, msg.ToNetType, msg.ToNodeID, et, msg.Pub[:], msg.Sig[:])
}

func (msg *MissingChainEventMsg) String() string {
	if msg == nil {
		return "MissChain<nil>"
	}
	return fmt.Sprintf("MissChain{ID:%s}", msg.ID)
}

func (msg *SevereErrorEventMsg) String() string {
	if msg == nil {
		return "SevereErr<nil>"
	}
	return fmt.Sprintf("SevereErr{ChainID:%s Err:%s}", msg.ChainID, msg.Err)
}

const (
	StartMsg = "start"
	StopMsg  = "stop"
)

type StartEMessage struct {
	Timestamp int64
	Signature []byte
}

type StopEMessage struct {
	Timestamp int64
	Signature []byte
}

func startOrStopTime(timestamp int64) error {
	local := time.Now().Unix()
	delta := int64(0)
	if local < timestamp {
		delta = timestamp - local
	} else {
		delta = local - timestamp
	}
	if delta > 30 {
		// an start message not in [local-30, local+30]
		return fmt.Errorf("invalid timestamp, now:%d timestamp:%d", local, timestamp)
	}
	return nil
}

func startOrStopDigest(act string, timestamp int64) []byte {
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, uint64(timestamp))
	return common.SystemHash256([]byte(act), ts)
}

func (m *StartEMessage) GetChainID() common.ChainID {
	return common.MainChainID
}

func (m *StartEMessage) Verify() error {
	if m == nil || len(m.Signature) != common.RealCipher.LengthOfSignature() {
		return errors.New("wrong size property")
	}
	if err := startOrStopTime(m.Timestamp); err != nil {
		return err
	}
	digest := startOrStopDigest(StartMsg, m.Timestamp)
	if !common.RealCipher.Verify(config.SystemStarterPK, digest, m.Signature) {
		return errors.New("signature verify failed")
	}
	return nil
}

func (m *StartEMessage) String() string {
	if m == nil {
		return "Start<nil>"
	}
	return fmt.Sprintf("Start{Timestamp:%d Sig:%x}", m.Timestamp, common.ForPrint(m.Signature))
}

func CreateStartMessage() (*StartEMessage, error) {
	if config.SystemStarterPrivate == nil {
		return nil, errors.New("not a starter")
	}
	timestamp := time.Now().Unix()
	digest := startOrStopDigest(StartMsg, timestamp)
	sig, err := common.RealCipher.Sign(config.SystemStarterPrivate.ToBytes(), digest)
	if err != nil {
		return nil, fmt.Errorf("sign the digest %x of start message failed: %v", digest, err)
	}
	return &StartEMessage{Timestamp: timestamp, Signature: sig}, nil
}

func (m *StopEMessage) GetChainID() common.ChainID {
	return common.MainChainID
}

func (m *StopEMessage) Verify() error {
	if m == nil || len(m.Signature) != common.RealCipher.LengthOfSignature() {
		return errors.New("wrong size property")
	}
	if err := startOrStopTime(m.Timestamp); err != nil {
		return err
	}
	digest := startOrStopDigest(StopMsg, m.Timestamp)
	if !common.RealCipher.Verify(config.SystemStarterPK, digest, m.Signature) {
		return errors.New("signature verify failed")
	}
	return nil
}

func (m *StopEMessage) String() string {
	if m == nil {
		return "Stop<nil>"
	}
	return fmt.Sprintf("Stop{Timestamp:%d Sig:%x)", m.Timestamp, common.ForPrint(m.Signature))
}

func CreateStopMessage() (*StopEMessage, error) {
	if config.SystemStarterPrivate == nil {
		return nil, errors.New("not a starter")
	}
	timestamp := time.Now().Unix()
	digest := startOrStopDigest(StopMsg, timestamp)
	sig, err := common.RealCipher.Sign(config.SystemStarterPrivate.ToBytes(), digest)
	if err != nil {
		return nil, fmt.Errorf("sign the digest %x of stop message failed: %v", digest, err)
	}
	return &StopEMessage{Timestamp: timestamp, Signature: sig}, nil
}
