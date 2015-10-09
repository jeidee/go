// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/jeidee/go/go-net"
	"github.com/jeidee/go/go-net/example/chat_test/chat_packet"
)

type ChatClient struct {
	*gonet.Client
	protocol gonet.Protocol

	stub  chat_packet.ChatStub
	proxy chat_packet.ChatProxy

	echoCnt int
	isLogin bool
}

func NewChatClient() *ChatClient {
	c := new(ChatClient)
	c.Client = new(gonet.Client)

	// Attach a packet protocol
	c.protocol = gonet.NewPacketProtocol()

	// Attach stub to handle incoming packets from server
	c.stub.OnResLogin = c.OnResLogin

	c.echoCnt = 0
	c.isLogin = false

	c.Init(nil, c, c.protocol)

	return c
}

/* public properties */
func (c *ChatClient) IsLogin() bool {
	return c.isLogin
}

/* ChatStub Event Handler */

func (c *ChatClient) OnResLogin(client *gonet.Client, packet *chat_packet.ResLogin) {
	fmt.Println("OnResLogin ----- ", packet.SenderId, packet.Result, c.echoCnt)

	if packet.Result == chat_packet.RESULT_OK {
		c.isLogin = true
	} else {
		c.isLogin = false
	}
}

/* Network Event Handler */

func (c *ChatClient) OnConnect(client *gonet.Client) {
	//fmt.Println("Client] OnConnect ", client)
}

func (c *ChatClient) OnClose(client *gonet.Client) {
	//fmt.Println("Client] OnClose ", client)
}

func (c *ChatClient) OnRead(client *gonet.Client, data interface{}) {
	packet := data.(*gonet.Packet)
	if packet == nil {
		fmt.Errorf("Invalid packet. %v\n", data)
		return
	}

	switch packet.GetId() {
	case chat_packet.PID_RES_LOGIN:
		obj := chat_packet.NewResLogin()
		gonet.PacketToObject(packet, &obj)
		c.stub.ResLogin(client, &obj)
		c.echoCnt++
	default:
		fmt.Println("Client] OnRead ...", packet.GetId())
	}
}

func (c *ChatClient) OnWrite(client *gonet.Client, data interface{}) {
	//msg := string(packet.GetBuffer()[:])
	//fmt.Println("Client] OnWrite ", client, msg)
}

func (c *ChatClient) OnError(client *gonet.Client, err error) {
	fmt.Println("Client] OnError ", client, err)
}
