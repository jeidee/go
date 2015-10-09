// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	_ "fmt"
	"github.com/jeidee/go/go-net/example/chat_test/chat_packet"
	"testing"
	"time"
)

const (
	SERVER_PORT = 9999
)

func newServer() *ChatServer {
	server := NewChatServer(SERVER_PORT)
	go server.Listen()

	return server
}

func TestLogin(t *testing.T) {
	server := newServer()

	client := NewChatClient()
	client.Connect("localhost", SERVER_PORT)

	var nickname chat_packet.NicknameT
	copy(nickname[:], "hello")

	client.proxy.ReqLogin(client.Client, nickname)
	time.Sleep(1 * time.Second)
	if !client.IsLogin() {
		t.Fatal("Login failed.")
	}

	if server.GetCu() != 1 {
		t.Fatal("Invalid the number of login users in server.")
	}
}

//func TestChat(t *testing.T) {
//	client := NewChatClient()
//	client.Connect("localhost", listenPort)

//	sampleCount := 100
//	for i := 0; i < sampleCount; i++ {
//		var nickname chat_packet.NicknameT
//		copy(nickname[:], "hello")
//		client.proxy.ReqLogin(client.Client, nickname)
//		client.echoCnt++
//	}

//	fmt.Println("echo count is ", client.echoCnt)
//}
