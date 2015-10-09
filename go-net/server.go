// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonet

import (
	"fmt"
	"net"
)

type EventChanType struct {
	Client *Client
	Data   interface{}
}

type Server struct {
	listener     net.Listener
	clients      map[net.Conn]*Client
	port         int16
	eventHandler NetworkEvent
	protocol     Protocol
	isRunning    bool

	// If you need to synchronize all network events,
	// you should use syncronous channel.
	useSyncChan   bool // If you need to use sync-channel, you should set true.
	OnReadChan    chan *EventChanType
	OnWriteChan   chan *EventChanType
	OnConnectChan chan *EventChanType
	OnCloseChan   chan *EventChanType
	OnErrorChan   chan *EventChanType
}

func NewServer(port int16, eventHandler NetworkEvent, protocol Protocol, useSyncChan bool) *Server {
	s := new(Server)

	s.Init(port, eventHandler, protocol, useSyncChan)

	return s
}

func (s *Server) Init(port int16, eventHandler NetworkEvent, protocol Protocol, useSyncChan bool) {
	s.protocol = protocol
	s.port = port
	s.eventHandler = eventHandler
	s.isRunning = false
	s.useSyncChan = useSyncChan

	if useSyncChan {
		s.OnReadChan = make(chan *EventChanType, 10)
		s.OnWriteChan = make(chan *EventChanType, 10)
		s.OnConnectChan = make(chan *EventChanType, 10)
		s.OnCloseChan = make(chan *EventChanType, 10)
		s.OnErrorChan = make(chan *EventChanType, 10)

		// Network events are proceeded goroutine-safely. 
		// If you need to synchronize shared objects of server,
		// you should write codes to this area.
		go func() {
			for {
				select {
				case data := <-s.OnConnectChan:
					s.eventHandler.OnConnect(data.Client)
				case data := <-s.OnCloseChan:
					s.eventHandler.OnClose(data.Client)
				case data := <-s.OnReadChan:
					s.eventHandler.OnRead(data.Client, data.Data)
				case data := <-s.OnWriteChan:
					s.eventHandler.OnWrite(data.Client, data.Data)
				case data := <-s.OnErrorChan:
					s.eventHandler.OnError(data.Client, data.Data.(error))
				}
			}
		}()
	}
}

func (s *Server) Panic(err error) {
	s.OnError(nil, err)
	s.Close()
}

func (s *Server) Listen() {
	if s.isRunning {
		s.OnError(nil, fmt.Errorf("Server is already listening."))
		return
	}

	s.clients = make(map[net.Conn]*Client)

	listenPort := fmt.Sprintf(":%d", s.port)
	var err error
	s.listener, err = net.Listen("tcp", listenPort)
	if err != nil {
		s.Panic(err)
		return
	}

	s.isRunning = true

	for {
		if s.isRunning == false {
			break
		}

		var conn net.Conn
		conn, err = s.listener.Accept()
		if err != nil {
			s.Panic(err)
			return
		}

		client := NewClient(conn, s, s.protocol)
		s.OnConnect(client)
		client.Dispatch()
	}

	s.Close()
}

func (s *Server) Close() {
	s.listener.Close()
	s.isRunning = false
	s.OnClose(nil)
}

// Implementation NetworkEvent interfaces.
// If you use sync-channel, network events would be sent each sync-channel.
// - goroutine-safe.
// - Because sync-channel listener is waited in goroutine, 
//   if you need to change shared objects,
//   you should write new sync-channel.
// If you don't use sync-channel, network events would be callbacked to event handlers.
// - goroutine not-safe

func (s *Server) OnConnect(client *Client) {
	if s.useSyncChan && s.OnConnectChan != nil {
		s.OnConnectChan <- &EventChanType{client, nil}
	} else {
		s.eventHandler.OnConnect(client)
	}
}

func (s *Server) OnClose(client *Client) {
	if s.useSyncChan && s.OnCloseChan != nil {
		s.OnCloseChan <- &EventChanType{client, nil}
	} else {
		s.eventHandler.OnClose(client)
	}
}

func (s *Server) OnRead(client *Client, data interface{}) {
	if s.useSyncChan && s.OnReadChan != nil {
		s.OnReadChan <- &EventChanType{client, data}
	} else {
		s.eventHandler.OnRead(client, data)
	}
}

func (s *Server) OnWrite(client *Client, data interface{}) {
	if s.useSyncChan && s.OnWriteChan != nil {
		s.OnWriteChan <- &EventChanType{client, data}
	} else {
		s.eventHandler.OnWrite(client, data)
	}
}

func (s *Server) OnError(client *Client, err error) {
	if s.useSyncChan && s.OnErrorChan != nil {
		s.OnErrorChan <- &EventChanType{client, err}
	} else {
		s.eventHandler.OnError(client, err)
	}
}
