// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat_packet

import (
	"github.com/jeidee/go/go-net"
)

type ChatStub struct {
	OnReqLogin    func(client *gonet.Client, packet *ReqLogin)
	OnResLogin    func(client *gonet.Client, packet *ResLogin)
	OnReqUserList func(client *gonet.Client, packet *ReqUserList)
	OnResUserList func(client *gonet.Client, packet *ResUserList)
	OnReqSendChat func(client *gonet.Client, packet *ReqSendChat)
	OnNotifyChat  func(client *gonet.Client, packet *NotifyChat)
}

func (stub *ChatStub) ReqLogin(client *gonet.Client, packet *ReqLogin) {
	if stub.OnReqLogin == nil {
		return
	}

	stub.OnReqLogin(client, packet)
}

func (stub *ChatStub) ResLogin(client *gonet.Client, packet *ResLogin) {
	if stub.OnResLogin == nil {
		return
	}

	stub.OnResLogin(client, packet)
}

func (stub *ChatStub) ReqUserList(client *gonet.Client, packet *ReqUserList) {
	if stub.OnReqUserList == nil {
		return
	}

	stub.OnReqUserList(client, packet)
}

func (stub *ChatStub) ResUserList(client *gonet.Client, packet *ResUserList) {
	if stub.OnResUserList == nil {
		return
	}

	stub.OnResUserList(client, packet)
}

func (stub *ChatStub) ReqSendChat(client *gonet.Client, packet *ReqSendChat) {
	if stub.OnReqSendChat == nil {
		return
	}

	stub.OnReqSendChat(client, packet)
}

func (stub *ChatStub) NotifyChat(client *gonet.Client, packet *NotifyChat) {
	if stub.OnNotifyChat == nil {
		return
	}

	stub.OnNotifyChat(client, packet)
}
