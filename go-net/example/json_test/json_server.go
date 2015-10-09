// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/jeidee/go/go-net"
)

type JsonServer struct {
	gonet.Server
	protocol gonet.Protocol
}

func NewJsonServer(port int16) *JsonServer {
	s := new(JsonServer)

	s.protocol = gonet.NewJsonProtocol()

	s.Init(port, s, s.protocol, true)

	return s
}

func (m *JsonServer) OnConnect(client *gonet.Client) {
	//fmt.Println("Server] OnConnect ", client)
}

func (m *JsonServer) OnClose(client *gonet.Client) {
	//fmt.Println("Server] OnClose ", client)
}

func (m *JsonServer) OnRead(client *gonet.Client, data interface{}) {
	//msg := string(packet.GetBuffer()[:])
	// echo
	client.Write(data)
	// log
	//fmt.Println("Server] OnRead ", client, msg)
}

func (m *JsonServer) OnWrite(client *gonet.Client, data interface{}) {
	//msg := string(packet.GetBuffer()[:])
	//fmt.Println("Server] OnWrite ", client, msg)
}

func (m *JsonServer) OnError(client *gonet.Client, err error) {
	fmt.Println("Server] OnError ", client, err)
}
