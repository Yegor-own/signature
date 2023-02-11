package main

import (
	"fmt"
	"net"
	"strings"
)

func process(conns map[int]net.Conn, n int) {
	var clientNum int

	conn := conns[n]
	defer conn.Close()
	for {
		buf := make([]byte, 256)
		readedLn, err := conns[n].Read(buf)
		if err != nil {
			if err.Error() == "EOF" {

				fmt.Println("Close", n)
				delete(conns, n)
			}
			fmt.Println(err)
		}

		message := string(buf[:readedLn])

		ln, err := fmt.Sscanf(message, "%d", &clientNum)
		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}
		if ln > 256 {
			fmt.Println("msg is too long, limit is 256")
		}

		pos := strings.Index(message, " ")
		if pos > 0 {
			outMsg := message[pos+1:]
			conn = conns[clientNum]
			if conn == nil {
				conns[n].Write([]byte("Connection is closed"))
				continue
			}
			outBuf := []byte(fmt.Sprintf("%d >> %s\n", clientNum, outMsg))
			ln, err = conn.Write(outBuf)
			if err != nil {
				fmt.Println("Server: ", err)
				break
			}
			if ln > 256 {
				fmt.Println("line is too long, limit is 256")
			}
		}
	}
}

func main() {
	fmt.Println("Server starting...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}

	conns := make(map[int]net.Conn, 1024)
	i := 0

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		conns[i] = conn
		go process(conns, i)
		i++
	}

}
