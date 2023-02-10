package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func readSock(conn net.Conn) {
	for {
		buf := make([]byte, 256)
		readedLn, _ := conn.Read(buf)
		if readedLn > 0 {
			fmt.Println(string(buf))
		}
	}
}

func readConsole(ch chan string) {
	for {
		fmt.Println(">")
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		out := line[:len(line)-1]
		ch <- out
	}
}

func main() {
	ch := make(chan string)
	defer close(ch)

	conn, _ := net.Dial("tcp", "localhost:8081")

	go readConsole(ch)
	go readSock(conn)

	for {
		val, ok := <-ch
		if ok {
			_, err := conn.Write([]byte(val))
			if err != nil {
				fmt.Println(err)
				break
			} else {
				time.Sleep(2 * time.Second)
			}
		}
	}

	fmt.Println("Finished...")
	conn.Close()
}
