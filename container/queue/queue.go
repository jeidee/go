// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This queue is not safe multi-goroutine!

package queue

import (
	"container/list"
)

type DefQueue struct {
	l *list.List
}

func NewDefQueue() *DefQueue{
	return &DefQueue{l: list.New()}
}

func (q *DefQueue) Push(v interface{}) {
	q.l.PushBack(v)
}

func (q *DefQueue) Pop() (interface{}, bool) {
	front := q.l.Front()
	v := front.Value
	q.l.Remove(front)
	
	return v, true
}

func (q *DefQueue) Len() int {
	return int(q.l.Len())
}