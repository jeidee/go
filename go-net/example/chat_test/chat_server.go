// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"github.com/jeidee/go/container/queue"
	"github.com/jeidee/go/go-net"
	"github.com/jeidee/go/go-net/example/chat_test/chat_packet"
)

type ChatUser struct {
	client    *gonet.Client
	sessionId int32
	nickname  string
}

const (
	MAX_USERS = 1000
)

type ChatServer struct {
	gonet.Server

	protocol      gonet.Protocol
	users         map[*gonet.Client]*ChatUser
	freeSessionId *queue.DefQueue
	stub          chat_packet.ChatStub
	proxy         chat_packet.ChatProxy
}

func NewChatServer(port int16) *ChatServer {
	s := new(ChatServer)

	s.protocol = gonet.NewPacketProtocol()

	s.stub.OnReqLogin = s.OnReqLogin
	s.users = make(map[*gonet.Client]*ChatUser)
	s.freeSessionId = queue.NewDefQueue()

	// Generate a pool for free session ids
	for i := 0; i < MAX_USERS; i++ {
		s.freeSessionId.Push(int32(i))
	}

	s.Init(port, s, s.protocol, true)

	return s
}

/* Public functions */

func (s *ChatServer) GetCu() int {
	return len(s.users)
}

/* Private functions */
//func (s *ChatServer) addUser(client *gonet.Client, nickname string) bool {
//	if s.users[client] == nil {
//		s.OnError(client, fmt.Errorf("Unregistered a client."))
//		return false
//	}

//	s.users[client].nickname = nickname
//	fmt.Printf("%s is logined.\n", s.users[client].nickname)

//}

/* ChatStub Event Handler */

func (s *ChatServer) OnReqLogin(client *gonet.Client, packet *chat_packet.ReqLogin) {
	nick := string(bytes.Trim(packet.Nickname[:], "\x00")[:])

	if s.users[client] == nil {
		s.OnError(client, fmt.Errorf("Unregistered a client"))
		return
	}

	s.users[client].nickname = nick
	fmt.Printf("%s is logined.\n", s.users[client].nickname)

	// Send success packt to server
	err := s.proxy.ResLogin(client, chat_packet.UserIdT(s.users[client].sessionId), chat_packet.RESULT_OK)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

/* Network Event Handler  */

func (s *ChatServer) OnConnect(client *gonet.Client) {
	user := ChatUser{client: client}
	v, _ := s.freeSessionId.Pop()
	user.sessionId = v.(int32)
	s.users[client] = &user

	fmt.Println("Server] OnConnect ", client)
}

func (s *ChatServer) OnClose(client *gonet.Client) {
	delete(s.users, client)

	fmt.Println("Server] OnClose ", client)
}

func (s *ChatServer) OnRead(client *gonet.Client, data interface{}) {
	fmt.Println("Server] OnRead ...", data)
	packet := data.(*gonet.Packet)
	if packet == nil {
		fmt.Errorf("Invalid packet. %v\n", data)
		return
	}
	switch packet.GetId() {
	case chat_packet.PID_REQ_LOGIN:
		obj := chat_packet.NewReqLogin()
		gonet.PacketToObject(packet, &obj)
		s.stub.ReqLogin(client, &obj)
	default:
		fmt.Println("Server] OnRead ...", packet.GetId())
	}
}

func (s *ChatServer) OnWrite(client *gonet.Client, data interface{}) {
	//msg := string(packet.GetBuffer()[:])
	//fmt.Println("Server] OnWrite ", client, msg)
}

func (s *ChatServer) OnError(client *gonet.Client, err error) {
	fmt.Println("Server] OnError ", client, err)
}
