// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queue

import (
	"sync/atomic"
)

type ChanQueue struct {
	syncChan chan interface{}
	size     int32
}

func NewChanQueue(maxSize int) *ChanQueue {
	return &ChanQueue{syncChan: make(chan interface{}, maxSize)}
}

func (q *ChanQueue) Push(v interface{}) {
	atomic.AddInt32(&q.size, 1)
	q.syncChan <- v
}

func (q *ChanQueue) Pop() (interface{}, bool) {
	atomic.AddInt32(&q.size, -1)
	v := <-q.syncChan
	return v, true
}

func (q *ChanQueue) Len() int {
	return int(q.size)
}