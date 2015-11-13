// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonet

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
)

type PacketIdT int32

type Packet struct {
	size int32
	id   PacketIdT
	buf  []byte
}

func NewPacket(buf []byte, size int32, id PacketIdT) *Packet {
	p := new(Packet)
	p.buf = buf
	p.size = size
	p.id = id

	return p
}

func NewPacketByObject(v interface{}, id PacketIdT) *Packet {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, v)

	p := new(Packet)
	p.buf = buf.Bytes()
	p.size = int32(buf.Len())
	p.id = id

	return p
}

func NewPacketByProto(msg proto.Message, id PacketIdT) *Packet {

	data, _ := proto.Marshal(msg)

	p := new(Packet)
	p.buf = data
	p.size = int32(len(data))
	p.id = id

	return p
}

func PacketToObject(packet *Packet, v interface{}) {
	buf := bytes.NewReader(packet.buf)

	binary.Read(buf, binary.BigEndian, v)
}

func PacketToProto(packet *Packet, msg proto.Message) {
	proto.Unmarshal(packet.buf, msg)
}

func (p *Packet) GetBuffer() []byte {
	return p.buf
}

func (p *Packet) GetSize() int32 {
	return p.size
}

func (p *Packet) GetId() PacketIdT {
	return p.id
}
