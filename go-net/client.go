// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonet

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	conn         net.Conn
	reader       *bufio.Reader
	eventHandler NetworkEvent
	protocol     Protocol
}

func NewClient(conn net.Conn, eventHandler NetworkEvent, protocol Protocol) *Client {
	c := new(Client)

	c.Init(conn, eventHandler, protocol)

	return c
}

func (c *Client) Init(conn net.Conn, eventHandler NetworkEvent, protocol Protocol) {
	c.conn = conn
	c.protocol = protocol
	c.reader = bufio.NewReader(c.conn)
	c.eventHandler = eventHandler
}

func (c *Client) Panic(err error) {
	c.eventHandler.OnError(c, err)
	c.Close()
}

func (c *Client) Connect(ip string, port int16) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	var err error
	c.conn, err = net.Dial("tcp", address)
	if err != nil {
		return false
	}

	c.eventHandler.OnConnect(c)

	c.reader = bufio.NewReader(c.conn)

	go c.Dispatch()

	return true
}

func (c *Client) Write(data interface{}) bool {
	_, err := c.protocol.Encode(c, data)
	if err != nil {
		c.Panic(err)
		return false
	}

	return true
}

func (c *Client) DispatchRead() {
	for {
		_, err := c.protocol.Decode(c)

		if err != nil {
			c.Panic(err)
		}
	}
}

func (c *Client) Dispatch() {
	go c.DispatchRead()
}

func (c *Client) Close() {
	c.conn.Close()
	c.eventHandler.OnClose(c)
}
