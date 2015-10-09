// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	server := NewJsonServer(9999)
	go server.Listen()

	os.Exit(m.Run())
}

func TestEcho(t *testing.T) {
	client := NewJsonClient()
	client.Connect("localhost", 9999)
	client.echobackChan = make(chan bool)

	testCount := 0
	sampleCount := 5000
	for i := 0; i < sampleCount; i++ {
		client.TestEchoBack("hello")
		ret := <-client.echobackChan
		if ret {
			testCount++
		}
	}
	
	if testCount != sampleCount {
		t.Error("Benchmark test failed!")
	}
}


func BenchmarkEcho(b *testing.B) {
	client := NewJsonClient()
	client.Connect("localhost", 9999)
	client.echobackChan = make(chan bool)

	testCount := 0
	for i := 0; i < b.N; i++ {
		client.TestEchoBack("hello")
		ret := <-client.echobackChan
		if ret {
			testCount++
		}
	}
	
	if testCount != b.N {
		b.Error("Benchmark test failed!")
	}
}
