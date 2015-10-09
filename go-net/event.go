// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonet

type NetworkEvent interface {
	OnConnect(client *Client)
	OnClose(client *Client)
	OnRead(client *Client, data interface{})
	OnWrite(client *Client, data interface{})
	OnError(client *Client, err error)
}
