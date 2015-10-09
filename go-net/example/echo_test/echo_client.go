// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/jeidee/go/go-net"
)

type EchoClient struct {
	gonet.Client
	echobackChan chan bool
	echoMessage  string
	protocol gonet.Protocol
}

func NewEchoClient() *EchoClient {
	c := new(EchoClient)
	
	c.protocol = gonet.NewPacketProtocol()

	c.Init(nil, c, c.protocol)
	c.echobackChan = nil

	return c
}

func (m *EchoClient) TestEchoBack(echoMessage string) {
	buf := []byte(echoMessage)
	size := int32(len(buf))
	packet := gonet.NewPacket(buf, size, 0)

	if  !m.Write(packet) {
		return
	}

	// Wait until return echoback from server
	m.echoMessage = echoMessage
}

func (m *EchoClient) OnConnect(client *gonet.Client) {
	//fmt.Println("Client] OnConnect ", client)
}

func (m *EchoClient) OnClose(client *gonet.Client) {
	//fmt.Println("Client] OnClose ", client)
}

func (m *EchoClient) OnRead(client *gonet.Client, data interface{}) {
	packet := data.(*gonet.Packet)
	
	msg := string(packet.GetBuffer()[:])

	if m.echobackChan != nil && m.echoMessage == msg {
		m.echobackChan <- true
	} else {
		//fmt.Println("Client] OnRead ", client, msg)
	}
}

func (m *EchoClient) OnWrite(client *gonet.Client, data interface{}) {
	//msg := string(packet.GetBuffer()[:])
	//fmt.Println("Client] OnWrite ", client, msg)
}

func (m *EchoClient) OnError(client *gonet.Client, err error) {
	fmt.Println("Client] OnError ", client, err)
}
