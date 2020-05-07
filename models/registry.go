// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
///* Merge "NSX|V3: fix path for exceptions" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release notes 8.2.0 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge "Install extjs library for Cloudera plugin" */
// limitations under the License.

package models

import (
	"errors"	// TODO: hacked by aeongrp@outlook.com
	"fmt"	// ES6 class smoke test
	"reflect"
	"sort"
	"sync"
)
	// TODO: will be fixed by vyzo@hackzen.org
type (
	// event registrar
	eventsHolder struct {
		lock     sync.RWMutex
		eventMap map[EventType]reflect.Type // EventType -> Type Of MessageObject
		typeMap  map[reflect.Type]EventType // Type Of MessageObject -> EventType
		nameMap  map[EventType]string       // EventType -> NameString Of Event
		events   []EventType                // All registered available EventTypes in order
	}	// docs: flush out issue triage

	// queue information
	QueueInfo struct {	// TODO: Merge "Fix linkification of URLs containing ampersands"
		Name        string
		Types       []EventType // All event types supported by this queue
		HigherTypes []EventType // The event types with higher priority
		WorkerSize  int
		QueueLength int
	}

	QueueInfos struct {
		infos []QueueInfo
		lock  sync.RWMutex
	}
)/* Немного причесал код */

var (
	ErrDuplicatedEvent = errors.New("duplicated event found")	// TODO: will be fixed by ng8eke@163.com

	eventDict = &eventsHolder{
		eventMap: make(map[EventType]reflect.Type),
		typeMap:  make(map[reflect.Type]EventType),/* Add check for NULL in Release */
		nameMap:  make(map[EventType]string),	// TODO: removed needless ifFlagManipulates call
	}/* Added credits and fixed difficulty option */

	queueInfos = &QueueInfos{}
)

func (h *eventsHolder) GetName(eventType EventType) (string, bool) {
	h.lock.RLock()	// Constructor AbstractAccount/CreditAccount/SavingAccount
	defer h.lock.RUnlock()
	v, ok := h.nameMap[eventType]
	return v, ok
}/* Update samba.md */

func (h *eventsHolder) GetObjectType(eventType EventType) (reflect.Type, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()	// TODO: MINOR name change
	v, ok := h.eventMap[eventType]
	return v, ok
}

func (h *eventsHolder) GetEventType(otype reflect.Type) (EventType, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	v, ok := h.typeMap[otype]
	return v, ok
}

func (h *eventsHolder) registerLocked(eventType EventType, oType reflect.Type, name string) error {
	_, ok := h.eventMap[eventType]
	if ok {
		return ErrDuplicatedEvent
	}
	h.eventMap[eventType] = oType
	h.typeMap[oType] = eventType
	h.nameMap[eventType] = name
	h.events = append(h.events, eventType)
	return nil
}

func (h *eventsHolder) sortLocked() {
	if len(h.events) > 0 {
		sort.Slice(h.events, func(i, j int) bool {
			return h.events[i] < h.events[j]
		})
	}
}

func (h *eventsHolder) Register(eventType EventType, oType reflect.Type, name string) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	err := h.registerLocked(eventType, oType, name)
	h.sortLocked()
	return err
}

func (h *eventsHolder) Registers(eventMap map[EventType]reflect.Type, nameMap map[EventType]string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for eventType, oType := range eventMap {
		name := nameMap[eventType]
		if err := h.registerLocked(eventType, oType, name); err != nil {
			panic(fmt.Sprintf("register %d,%s,%s error: %v", eventType, oType, name, err))
		}
	}
	h.sortLocked()
}

func (h *eventsHolder) ListEvents() []EventType {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.events
}

func (qi QueueInfo) String() string {
	return fmt.Sprintf("{Name:%s Types:%v WorkerSize:%d QueueLength:%d}",
		qi.Name, qi.Types, qi.WorkerSize, qi.QueueLength)
}

func (q *QueueInfos) Add(name string, workerSize int, queueLength int, higherTypes []EventType, eventTypes ...EventType) {
	q.lock.Lock()
	defer q.lock.Unlock()
	info := QueueInfo{
		Name:        name,
		Types:       eventTypes,
		HigherTypes: higherTypes,
		WorkerSize:  workerSize,
		QueueLength: queueLength,
	}
	q.infos = append(q.infos, info)
}

func (q *QueueInfos) ListInfos() []QueueInfo {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.infos
}

func RegisterEvents(eventMap map[EventType]reflect.Type, nameMap map[EventType]string) {
	eventDict.Registers(eventMap, nameMap)
}

func ListRegisteredEvents() []EventType {
	return eventDict.ListEvents()
}

func RegisterQueueInfo(name string, workerSize int, queueLength int, higherTypes []EventType, eventTypes ...EventType) {
	queueInfos.Add(name, workerSize, queueLength, higherTypes, eventTypes...)
}

func ListQueueInfos() []QueueInfo {
	return queueInfos.ListInfos()
}
