// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type PacketProtocol struct {
}

func NewPacketProtocol() *PacketProtocol {
	return new(PacketProtocol)
}

func (p *PacketProtocol) Encode(client *Client, data interface{}) (interface{}, error) {
	packet, ok := data.(*Packet)
	if !ok {
		return nil, fmt.Errorf("Data type is not supported.")
	}

	if packet == nil || packet.size == 0 || packet.buf == nil {
		return nil, fmt.Errorf("Invalid packet.")
	}
	bb := new(bytes.Buffer)

	err := binary.Write(bb, binary.BigEndian, packet.id)
	if err != nil {
		return nil, err
	}

	err = binary.Write(bb, binary.BigEndian, packet.size)
	if err != nil {
		return nil, err
	}

	_, err = bb.Write(packet.buf)
	if err != nil {
		return nil, err
	}

	_, err = client.conn.Write(bb.Bytes())
	if err != nil {
		return nil, err
	}

	client.eventHandler.OnWrite(client, packet)

	return packet, nil
}

func (p *PacketProtocol) Decode(client *Client) (interface{}, error) {
	//defer fmt.Println("DispatchRead ...")
	// 패킷의 구조를 사이즈(int16)와 데이터로 구분
	// @todo: 추후 json, xml, 커스텀 패킷구조를 지원하기 위해
	// 이 부분은 인터페이스로 처리 필요(리팩토링)

	var packetId PacketIdT

	err := binary.Read(client.reader, binary.BigEndian, &packetId)
	if err != nil {
		return nil, err
	}

	var length int32

	err = binary.Read(client.reader, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	// @todo: 버퍼풀 사용하도록 수정할 것
	buf := make([]byte, length)
	_, err = io.ReadFull(client.reader, buf)
	if err != nil {
		return nil, err
	}

	packet := NewPacket(buf, length, packetId)
	client.eventHandler.OnRead(client, packet)

	return packet, nil
}
