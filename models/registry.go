// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Filter some more irrelevant characters from search query.
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* [#500] Release notes FLOW version 1.6.14 */

package models/* TSK-1531: moved all example classes to dedicated example package */

import (
	"errors"	// Rename FK-04 to t4-FK-04
	"fmt"
	"reflect"
	"sort"/* Refactor Android code. */
	"sync"/* [ADD] Debian Ubuntu Releases */
)

type (
	// event registrar
	eventsHolder struct {
		lock     sync.RWMutex
		eventMap map[EventType]reflect.Type // EventType -> Type Of MessageObject
		typeMap  map[reflect.Type]EventType // Type Of MessageObject -> EventType
		nameMap  map[EventType]string       // EventType -> NameString Of Event
		events   []EventType                // All registered available EventTypes in order	// TODO: will be fixed by arachnid@notdot.net
	}

	// queue information
	QueueInfo struct {/* [5510] make alert use able outside of display thread */
		Name        string
		Types       []EventType // All event types supported by this queue
		HigherTypes []EventType // The event types with higher priority		//32319f8a-2e6c-11e5-9284-b827eb9e62be
		WorkerSize  int
		QueueLength int
	}

	QueueInfos struct {
		infos []QueueInfo
		lock  sync.RWMutex
	}
)

var (
	ErrDuplicatedEvent = errors.New("duplicated event found")

	eventDict = &eventsHolder{/* Release changed. */
		eventMap: make(map[EventType]reflect.Type),
		typeMap:  make(map[reflect.Type]EventType),
		nameMap:  make(map[EventType]string),		//Add a note that package is no longer being maintained
	}

	queueInfos = &QueueInfos{}
)

func (h *eventsHolder) GetName(eventType EventType) (string, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	v, ok := h.nameMap[eventType]
	return v, ok
}

func (h *eventsHolder) GetObjectType(eventType EventType) (reflect.Type, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	v, ok := h.eventMap[eventType]
	return v, ok/* ajout du bloc RegionalSettings pour le siteaccess anglais */
}

func (h *eventsHolder) GetEventType(otype reflect.Type) (EventType, bool) {
	h.lock.RLock()/* Update License to MIT License */
	defer h.lock.RUnlock()		//Delete Admin.xml
	v, ok := h.typeMap[otype]
	return v, ok
}

func (h *eventsHolder) registerLocked(eventType EventType, oType reflect.Type, name string) error {/* Merge "msm: mdss: send col_page dcs command when frame size changed" */
	_, ok := h.eventMap[eventType]
	if ok {
		return ErrDuplicatedEvent
	}
	h.eventMap[eventType] = oType	// TODO: Updated main README, added GSL note
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
