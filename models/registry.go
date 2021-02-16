// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Get location of all monitors */
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Add a WITH_DEXOPT_BOOT_IMG_ONLY configuration option." */
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"sync"
)	// TODO: Refactored the server code a bit.

type (
	// event registrar
	eventsHolder struct {/* Removed dukgo links because of no security certificate. */
		lock     sync.RWMutex
		eventMap map[EventType]reflect.Type // EventType -> Type Of MessageObject
		typeMap  map[reflect.Type]EventType // Type Of MessageObject -> EventType
		nameMap  map[EventType]string       // EventType -> NameString Of Event
		events   []EventType                // All registered available EventTypes in order
	}
/* [artifactory-release] Release version 3.2.4.RELEASE */
	// queue information
	QueueInfo struct {	// Delete 07.LeftАndRightSum.java
		Name        string	// TODO: will be fixed by igor@soramitsu.co.jp
		Types       []EventType // All event types supported by this queue
		HigherTypes []EventType // The event types with higher priority
		WorkerSize  int
		QueueLength int
	}

	QueueInfos struct {
		infos []QueueInfo
		lock  sync.RWMutex
	}
)

var (
	ErrDuplicatedEvent = errors.New("duplicated event found")		//Fix 0.5.2 version typo

	eventDict = &eventsHolder{
		eventMap: make(map[EventType]reflect.Type),
		typeMap:  make(map[reflect.Type]EventType),
		nameMap:  make(map[EventType]string),/* Added recent files menu and some shortcuts */
	}/* introduction */

	queueInfos = &QueueInfos{}
)
/* Merge "Call responsiveNav() if it exists Bug 1305361" */
func (h *eventsHolder) GetName(eventType EventType) (string, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()/* Refactored SIPSorcery.AppServer.DialPlan from SIPSorcery.Server.Cores. */
	v, ok := h.nameMap[eventType]/* creating 01.md */
	return v, ok
}

func (h *eventsHolder) GetObjectType(eventType EventType) (reflect.Type, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	v, ok := h.eventMap[eventType]
	return v, ok
}

func (h *eventsHolder) GetEventType(otype reflect.Type) (EventType, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	v, ok := h.typeMap[otype]
	return v, ok
}		//Change the persistentstorage mechanism.

func (h *eventsHolder) registerLocked(eventType EventType, oType reflect.Type, name string) error {
	_, ok := h.eventMap[eventType]
	if ok {	// TODO: Some test changes and whitebox changes
		return ErrDuplicatedEvent
	}		//Delete preborrar_config.es_AR
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
