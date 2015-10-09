// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat_packet

const (
	CHAT_PACKET_VERSION = 100
	PID_USER_INFO       = 101
	PID_ROOM_INFO       = 102
	PID_REQ_LOGIN       = 1001
	PID_RES_LOGIN       = 1002
	PID_REQ_USER_LIST   = 1003
	PID_RES_USER_LIST   = 1004
	PID_REQ_SEND_CHAT   = 1005
	PID_NOTIFY_CHAT     = 1006
)

const (
	RESULT_OK   = 0
	RESULT_FAIL = 1
)

type UserIdT int32
type ResultT int16
type ChatT [1024]byte
type NicknameT [24]byte
type TitleT [300]byte

type UserInfo struct {
	Id       UserIdT
	Nickname NicknameT
}

// Request Login.
type ReqLogin struct {
	Nickname NicknameT
}

func NewReqLogin() ReqLogin {
	return ReqLogin{}
}

// Response Login.
type ResLogin struct {
	SenderId UserIdT // Session id is issued by server when login succeed.
	Result   ResultT // RESULT_OK, RESULT_FAIL, ETC
}

func NewResLogin() ResLogin {
	return ResLogin{}
}

// Request user list
type ReqUserList struct {
	SenderId UserIdT
}

func NewReqUserList() ReqUserList {
	return ReqUserList{}
}

// Response user list
type ResUserList struct {
	Page          int16
	PageSize      int16
	NumberOfUsers int16
	Users         [10]UserInfo
}

func NewResUserList() ResUserList {
	return ResUserList{}
}

// Send chat message.
// Response isn't required.
type ReqSendChat struct {
	SenderId UserIdT
	Chat     ChatT // utf8 string, max length is 500 runes
}

func NewReqSendChat() ReqSendChat {
	return ReqSendChat{}
}

// Notify chat message to all users
type NotifyChat struct {
	SenderId UserIdT
	Chat     ChatT
}

func NewNotifyChat() NotifyChat {
	return NotifyChat{}
}
