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

func (proxy *ChatProxy) ReqLogin(client *gonet.Client, obj *ReqLogin) error {
	fmt.Println("ReqLogin...")
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	packet := gonet.NewPacketByProto(obj, gonet.PacketIdT(PacketType_REQ_LOGIN))

	if !client.Write(packet) {
		return fmt.Errorf("ReqLogin: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ResLogin(client *gonet.Client, obj *ResLogin) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	packet := gonet.NewPacketByProto(obj, gonet.PacketIdT(PacketType_RES_LOGIN))

	if !client.Write(packet) {
		return fmt.Errorf("ResLogin: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ReqUserList(client *gonet.Client, obj *ReqUserList) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	packet := gonet.NewPacketByProto(obj, gonet.PacketIdT(PacketType_REQ_USER_LIST))

	if !client.Write(packet) {
		return fmt.Errorf("ReqUserList: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ResUserList(client *gonet.Client, obj *ReqUserList) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	packet := gonet.NewPacketByProto(obj, gonet.PacketIdT(PacketType_REQ_USER_LIST))

	if !client.Write(packet) {
		return fmt.Errorf("ResUserList: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) ReqSendChat(client *gonet.Client, obj *ReqSendChat) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	packet := gonet.NewPacketByProto(obj, gonet.PacketIdT(PacketType_REQ_SEND_CHAT))

	if !client.Write(packet) {
		return fmt.Errorf("ReqSendChat: client.Write() fail.")
	}

	return nil
}

func (proxy *ChatProxy) NotifyChat(client *gonet.Client, obj *NotifyChat) error {
	if client == nil {
		return fmt.Errorf("gonet.Client is nil.")
	}

	packet := gonet.NewPacketByProto(obj, gonet.PacketIdT(PacketType_NOTIFY_CHAT))

	if !client.Write(packet) {
		return fmt.Errorf("NotifyChat : client.Write() fail.")
	}

	return nil
}
