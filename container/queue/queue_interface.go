// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package queue

type Queue interface {
	Push(v interface{})
	Pop() (interface{}, bool)
	Len() int
}