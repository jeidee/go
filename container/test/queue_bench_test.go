// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queue

import (
	"github.com/jeidee/go/container/queue"
	"testing"
)

func BenchmarkChanQueue(b *testing.B) {
	q := queue.NewChanQueue(b.N)
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}