// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/jeidee/go/go-net"
)

type EchoServer struct {
	gonet.Server
	protocol gonet.Protocol
}

func NewEchoServer(port int16) *EchoServer {
	s := new(EchoServer)

	s.protocol = gonet.NewPacketProtocol()

	s.Init(port, s, s.protocol, true)

	return s
}

func (m *EchoServer) OnConnect(client *gonet.Client) {
	//fmt.Println("Server] OnConnect ", client)
}

func (m *EchoServer) OnClose(client *gonet.Client) {
	//fmt.Println("Server] OnClose ", client)
}

func (m *EchoServer) OnRead(client *gonet.Client, data interface{}) {
	//msg := string(packet.GetBuffer()[:])
	// echo
	client.Write(data)
	// log
	//fmt.Println("Server] OnRead ", client, msg)
}

func (m *EchoServer) OnWrite(client *gonet.Client, data interface{}) {
	//msg := string(packet.GetBuffer()[:])
	//fmt.Println("Server] OnWrite ", client, msg)
}

func (m *EchoServer) OnError(client *gonet.Client, err error) {
	fmt.Println("Server] OnError ", client, err)
}
