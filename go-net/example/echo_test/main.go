// Copyright 2015 <jeidee@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jeidee/go/go-net"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	server := NewEchoServer(8888)
	go server.Listen()

	client := NewEchoClient()
	client.Connect("localhost", 8888)

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
		case "send":
			if len(tokens) != 2 {
				continue
			}
			buf := []byte(strings.Trim(tokens[1], " \r\n"))
			size := int32(len(buf))
			packet := gonet.NewPacket(buf, size, 0)
			_, err = client.protocol.Encode(&client.Client, packet)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
