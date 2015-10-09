// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonet

import (
	"encoding/json"
)

type JsonProtocol struct {
}

func NewJsonProtocol() *JsonProtocol {
	return new(JsonProtocol)
}

func (p *JsonProtocol) Encode(client *Client, data interface{}) (interface{}, error) {

	e := json.NewEncoder(client.conn)
	err := e.Encode(data)

	if err != nil {
		return nil, err
	}

	client.eventHandler.OnWrite(client, data)

	return data, err
}

func (p *JsonProtocol) Decode(client *Client) (interface{}, error) {
	d := json.NewDecoder(client.conn)

	var data interface{}
	err := d.Decode(&data)
	if err != nil {
		return nil, err
	}

	client.eventHandler.OnRead(client, data)

	return data, nil
}
