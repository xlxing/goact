package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// 监听端口
	ln, _ := net.Listen("tcp", ":8081")

	// 接收连接
	conn, _ := ln.Accept()

	// loop
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println("Message Received:", string(message))

		newmessage := strings.ToUpper(message)

		conn.Write([]byte(newmessage + "\n"))
	}
}
