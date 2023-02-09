package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server starting...")
	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message received:", string(message))
		conn.Write([]byte("Lutoe " + message + "\n"))
	}
}
