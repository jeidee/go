// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonet

type Protocol interface {
	Encode(client *Client, data interface{}) (interface{}, error)
	Decode(client *Client) (interface{}, error)
}
