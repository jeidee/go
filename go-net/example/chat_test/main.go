// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"github.com/jeidee/go/go-net/example/chat_test/chat_packet"
	"os"
	"strings"
)

func main() {
	server := NewChatServer(8888)
	go server.Listen()

	client := NewChatClient()
	if !client.Connect("localhost", 8888) {
		fmt.Errorf("Can't connect to server.")
		return
	}

	for {
		fmt.Println("Press quit to terminate process.")

		buffer := bufio.NewReader(os.Stdin)
		input, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println("Error ! : ", err)
			return
		}

		tokens := strings.Split(input, " ")
		cmd := strings.Trim(tokens[0], " \r\n")

		switch cmd {
		case "quit":
			server.Close()
			os.Exit(0)
		case "login":
			if len(tokens) != 2 {
				continue
			}
			var nick chat_packet.NicknameT
			copy(nick[:], strings.Trim(tokens[1], " \r\n"))
			client.proxy.ReqLogin(client.Client, nick)
		}
	}
}
