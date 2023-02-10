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

		_, err = fmt.Sscanf(message, "%d", &clientNum)
		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
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
			_, err = conn.Write(outBuf)
			if err != nil {
				fmt.Println("Server: ", err)
				break
			}
		}
	}
}

func main() {
	fmt.Println("Server starting...")
	ln, _ := net.Listen("tcp", ":8081")
	conns := make(map[int]net.Conn, 1024)
	i := 0

	for {
		conn, _ := ln.Accept()
		conns[i] = conn
		go process(conns, i)
		i++
	}

}
