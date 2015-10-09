// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queue

import (
	"github.com/jeidee/go/container/queue"
	"runtime"
	"sync"
	"testing"
)

// Test push and pop of queue using normal queue not-safe goroutine
func TestDefQueue(t *testing.T) {
	q := queue.NewDefQueue()

	q.Push("hello")
	q.Push("world")

	v, ok := q.Pop()

	if !ok {
		t.Error("Pop failed!")
	}

	if v.(string) != "hello" {
		t.Errorf("Expected : hello, Result : %s", v.(string))
	}

	if q.Len() != 1 {
		t.Errorf("Expected queue size : 1, Result : %d", q.Len())
	}

	v, ok = q.Pop()

	if !ok {
		t.Error("Pop failed!")
	}

	if v.(string) != "world" {
		t.Errorf("Expected : world, Result : %s", v.(string))
	}

	if q.Len() != 0 {
		t.Errorf("Expected queue size : 0, Result : %d", q.Len())
	}
}

// Test push and pop of queue using channel for syncronous queueing
func TestChanQueue(t *testing.T) {
	q := queue.NewChanQueue(2)

	q.Push("hello")
	q.Push("world")

	v, ok := q.Pop()

	if !ok {
		t.Error("Pop failed!")
	}

	if v.(string) != "hello" {
		t.Errorf("Expected : hello, Result : %s", v.(string))
	}

	if q.Len() != 1 {
		t.Errorf("Expected queue size : 1, Result : %d", q.Len())
	}

	v, ok = q.Pop()

	if !ok {
		t.Error("Pop failed!")
	}

	if v.(string) != "world" {
		t.Errorf("Expected : world, Result : %s", v.(string))
	}

	if q.Len() != 0 {
		t.Errorf("Expected queue size : 0, Result : %d", q.Len())
	}
}

// Test push and pop of channel queue using goroutine
func TestChanQueueSync(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	q := queue.NewChanQueue(10)

	wg := new(sync.WaitGroup)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			for j := (id * 100) + 0; j < (id*100)+100; j++ {
				q.Push(j)
			}
			wg.Done()
		}(i)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			for j := 0; j < 100; j++ {
				for {
					_, ok := q.Pop()
					if ok {
						//fmt.Printf("%d : %v\n", id, v)
						break
					}
				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	if q.Len() != 0 {
		t.Error("Queue size is : ", q.Len())
	}
}