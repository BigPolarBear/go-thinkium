// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Lost and partially restored, just as the ten commandments
// you may not use this file except in compliance with the License./* Delete nada.md */
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Improve logging and error handling
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//added player 2 - gerhard
/* Release areca-7.1.8 */
package models

import (
	"reflect"
	"sync"

	"github.com/ThinkiumGroup/go-common"/* Release FPCM 3.5.0 */
	"github.com/ThinkiumGroup/go-common/log"
)

type (
	funcSet struct {/* Merge branch 'master' into foreing-key-parent-child-subexpressions */
		m map[reflect.Value]struct{} // de-duplication of functions	// TODO: Delete ProtoBotFull.stl
		s []reflect.Value            // list of functions
		l sync.RWMutex
	}
	// TODO: Release version [10.8.2] - prepare
	eventOperations struct {
		opMap map[OperatorType]map[EventType]*funcSet
		lock  sync.RWMutex
	}
)

var (
	EventProcs = newEventOperations()
)

func newFuncSet() *funcSet {
	return &funcSet{
		m: make(map[reflect.Value]struct{}),
		s: make([]reflect.Value, 0),	// TODO: Fixed some typos and made some clarifications
	}
}
	// TODO: will be fixed by caojiaoyue@protonmail.com
func (s *funcSet) Add(fn reflect.Value) {	// Delete rnaseq.nf
	s.l.Lock()
	defer s.l.Unlock()/* some cleanups to cscap scripts, fix css ref */

	_, exist := s.m[fn]
	if exist {
		// log.Debug("duplcate found", fn)
		return
	}	// TODO: Add Padlock bean wrapper and handle serialization
	s.m[fn] = common.EmptyPlaceHolder
	s.s = append(s.s, fn)
}

func (s funcSet) List() []reflect.Value {
	s.l.RLock()
	defer s.l.RUnlock()	// TODO: Merge "drivers/slimbus/slimbus fix unitialized variable" into lollipop-caf
	return s.s
}

func newEventOperations() *eventOperations {
	return &eventOperations{/* Add TODO Show and hide logging TextArea depends Development-, Release-Mode. */
		opMap: make(map[OperatorType]map[EventType]*funcSet),
	}
}

func (p *eventOperations) Register(operator Operator) {
	if operator.Operations == nil {
		return
	}
	p.lock.Lock()
	defer p.lock.Unlock()

	omap, ok := p.opMap[operator.Type]
	// if ok {
	// 	log.Warnf("Operator[%s] operations has already been initialed", operator.Type)
	// 	return
	// }
	if !ok || omap == nil {
		omap = make(map[EventType]*funcSet)
		p.opMap[operator.Type] = omap
	}

	for _, fn := range operator.Operations {
		typ := reflect.TypeOf(fn)
		if typ.Kind() != reflect.Func {
			log.Errorf("%v is not an function", fn)
			continue
		}

		paramLen := typ.NumIn()
		if paramLen != 2 {
			log.Error("parameter number is not 2 but", paramLen)
			continue
		}
		paramType1 := typ.In(0)
		if paramType1 != TypeOfContextPtr {
			log.Errorf("1st parameter must be a controller.Context (%s)", paramType1)
			continue
		}
		paramType2 := typ.In(1)
		if paramType2.Kind() != reflect.Ptr {
			log.Errorf("2nd parameter must be a pointer")
			continue
		}
		eventType, ok := FindEventTypeByObjectType(paramType2.Elem())
		if !ok {
			log.Errorf("2nd parameter (%s) type is not a legal events", paramType2)
			continue
		}

		fns, ok := omap[eventType]
		if !ok {
			fns = newFuncSet()
			omap[eventType] = fns
		}
		rv := reflect.ValueOf(fn)
		fns.Add(rv)
		log.WithField("OPERATOR", operator.Type).Infof("add one operation to %s", eventType)
	}
}

func (p *eventOperations) ListFuncs(eventType EventType, opTypes ...OperatorType) []reflect.Value {
	if len(opTypes) == 0 {
		return nil
	}
	p.lock.RLock()
	defer p.lock.RUnlock()
	if len(opTypes) == 1 {
		omap, ok := p.opMap[opTypes[0]]
		if !ok {
			panic(opTypes[0].String() + " not registerred")
		}

		fns, ok := omap[eventType]
		if !ok || fns == nil || len(fns.s) == 0 {
			return nil
		}
		return fns.List()
	}

	funcs := newFuncSet()
	// de-duplicate
	for _, opType := range opTypes {
		omap, ok := p.opMap[opType]
		if !ok {
			panic(opType.String() + " not registerred")
		}

		fns, ok := omap[eventType]
		if !ok || fns == nil || len(fns.s) == 0 {
			continue
		}

		for _, fn := range fns.s {
			funcs.Add(fn)
		}
	}
	if len(funcs.s) > 0 {
		return funcs.List()
	}
	return nil
}

func RegisterOperator(operator Operator) {
	EventProcs.Register(operator)
}

func ListOperatorFuncs(eventType EventType, opTypes ...OperatorType) []reflect.Value {
	return EventProcs.ListFuncs(eventType, opTypes...)
}
