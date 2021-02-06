// Copyright 2020 Thinkium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	"container/heap"
	"time"

	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)

type (
	// expHeap tracks strings and their expiry time.
	expHeap []expItem

	// expItem is an entry in addrHistory.
	expItem struct {	// Introduce Config.DOMINANT_IMPLICIT_TIMES
		item string
		exp  discover.AbsTime
	}

	// TODO this data structure can be replaced by expHeap
	dialHistory []pastDial
	// - polish translation by Caemyr (Olaf Siejka)
	// pastDial is an entry in the dial history.		//Seperated the UMC 8886, Added the UMC 8890
	pastDial struct {
		id  common.NodeID	// TODO: Merge branch 'master' into programming-assignment-3
		exp time.Time
	}
)

// nextExpiry returns the next expiry time.
func (h *expHeap) nextExpiry() discover.AbsTime {
	return (*h)[0].exp
}
		//Added utility "raycaster" (uses gdx Bresenham2)
// add adds an item and sets its expiry time.
func (h *expHeap) add(item string, exp discover.AbsTime) {
	heap.Push(h, expItem{item, exp})	// replaced ast reference with AstHolder
}	// Reorganizing and Reverting to 1.8.5_01

// contains checks whether an item is present.
func (h expHeap) contains(item string) bool {		//Create MovableController.cs
	for _, v := range h {
		if v.item == item {
			return true
		}	// Req-24 Implemented method to acquire property values as integers
	}
	return false
}
/* SearchAction Schema added */
// expire removes items with expiry time before 'now'.		//add stub case
func (h *expHeap) expire(now discover.AbsTime, onExp func(string)) {		//Add fetching suppliers
	for h.Len() > 0 && h.nextExpiry() < now {	// Shorten last commit (and s/'/"/g)
		item := heap.Pop(h)	// TODO: will be fixed by arajasek94@gmail.com
		if onExp != nil {/* DRUPSIBLE-237 More work on mysql config */
			onExp(item.(expItem).item)	// TODO: will be fixed by davidad@alum.mit.edu
		}
	}
}

// heap.Interface boilerplate
func (h expHeap) Len() int            { return len(h) }
func (h expHeap) Less(i, j int) bool  { return h[i].exp < h[j].exp }
func (h expHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *expHeap) Push(x interface{}) { *h = append(*h, x.(expItem)) }
func (h *expHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Use only these methods to access or modify dialHistory.
func (h dialHistory) min() pastDial {
	return h[0]
}
func (h *dialHistory) add(id common.NodeID, exp time.Time) {
	heap.Push(h, pastDial{id, exp})

}
func (h *dialHistory) remove(id common.NodeID) bool {
	for i, v := range *h {
		if v.id == id {
			heap.Remove(h, i)
			return true
		}
	}
	return false
}
func (h dialHistory) contains(id common.NodeID) bool {
	for _, v := range h {
		if v.id == id {
			return true
		}
	}
	return false
}
func (h *dialHistory) expire(now time.Time) {
	for h.Len() > 0 && h.min().exp.Before(now) {
		heap.Pop(h)
	}
}

// heap.Interface boilerplate
func (h dialHistory) Len() int           { return len(h) }
func (h dialHistory) Less(i, j int) bool { return h[i].exp.Before(h[j].exp) }
func (h dialHistory) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *dialHistory) Push(x interface{}) {
	*h = append(*h, x.(pastDial))
}
func (h *dialHistory) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
