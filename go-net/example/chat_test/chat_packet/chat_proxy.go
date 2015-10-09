// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat_packet

import (
	"fmt"

	"github.com/jeidee/go/go-net"
)

type ChatProxy struct {
}

func (proxy *ChatProxy) ReqLogin(client *gonet.Client, nickname NicknameT) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	obj := NewReqLogin()
	obj.Nickname = nickname

	packet := gonet.NewPacketByObject(obj, PID_REQ_LOGIN)

	if !client.Write(packet) {
		return fmt.Errorf("ReqLogin: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ResLogin(client *gonet.Client, senderId UserIdT, result ResultT) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	obj := NewResLogin()
	obj.SenderId = senderId
	obj.Result = result

	packet := gonet.NewPacketByObject(obj, PID_RES_LOGIN)

	if !client.Write(packet) {
		return fmt.Errorf("ResLogin: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ReqUserList(client *gonet.Client, senderId UserIdT) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	obj := NewReqUserList()
	obj.SenderId = senderId

	packet := gonet.NewPacketByObject(obj, PID_REQ_USER_LIST)

	if !client.Write(packet) {
		return fmt.Errorf("ReqUserList: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ResUserList(client *gonet.Client, page int16, pageSize int16, numberOfUsers int16, users [10]UserInfo) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	obj := NewResUserList()
	obj.Page = page
	obj.PageSize = pageSize
	obj.NumberOfUsers = numberOfUsers
	obj.Users = users

	packet := gonet.NewPacketByObject(obj, PID_RES_USER_LIST)

	if !client.Write(packet) {
		return fmt.Errorf("ResUserList: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ReqSendChat(client *gonet.Client, senderId UserIdT, chat ChatT) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	obj := NewReqSendChat()
	obj.SenderId = senderId
	obj.Chat = chat

	packet := gonet.NewPacketByObject(obj, PID_REQ_SEND_CHAT)

	if !client.Write(packet) {
		return fmt.Errorf("ReqSendChat: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) NotifyChat(client *gonet.Client, senderId UserIdT, chat ChatT) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	obj := NewNotifyChat()
	obj.SenderId = senderId
	obj.Chat = chat

	packet := gonet.NewPacketByObject(obj, PID_NOTIFY_CHAT)

	if !client.Write(packet) {
		return fmt.Errorf("NotifyChat : client.Write() fail.")
	}

	return nil
}
