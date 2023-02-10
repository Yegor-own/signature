package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message received:", string(message))
		_, err := conn.Write([]byte("Lutoe " + message + "\n"))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func main() {
	fmt.Println("Server starting...")
	ln, _ := net.Listen("tcp", ":8081")

	for {
		conn, _ := ln.Accept()
		go process(conn)
	}

}
