// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/jeidee/go/go-net"
)

type JsonClient struct {
	gonet.Client
	echobackChan chan bool
	echoMessage  string
	protocol gonet.Protocol
}

func NewJsonClient() *JsonClient {
	c := new(JsonClient)
	
	c.protocol = gonet.NewJsonProtocol()

	c.Init(nil, c, c.protocol)
	c.echobackChan = nil

	return c
}

func (m *JsonClient) TestEchoBack(echoMessage string) {
	msg := make(map[string]interface{})
	msg["msg"] = echoMessage

	if  !m.Write(msg) {
		return
	}

	// Wait until return echoback from server
	m.echoMessage = echoMessage
}

func (m *JsonClient) OnConnect(client *gonet.Client) {
	//fmt.Println("Client] OnConnect ", client)
}

func (m *JsonClient) OnClose(client *gonet.Client) {
	//fmt.Println("Client] OnClose ", client)
}

func (m *JsonClient) OnRead(client *gonet.Client, data interface{}) {
	msg := data.(map[string]interface{})
	
	if m.echobackChan != nil && m.echoMessage == msg["msg"] {
		m.echobackChan <- true
	} else {
		fmt.Println("Client] OnRead ", client, msg)
	}
}

func (m *JsonClient) OnWrite(client *gonet.Client, data interface{}) {
	//msg := data.(map[string]interface{})
	//fmt.Println("Client] OnWrite ", client, msg)
}

func (m *JsonClient) OnError(client *gonet.Client, err error) {
	fmt.Println("Client] OnError ", client, err)
}
