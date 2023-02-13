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
		readedLn, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		if readedLn > 0 {
			fmt.Println(string(buf))
		}
	}
}

func readConsole(ch chan string) {
	for {
		fmt.Println(">")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if len(line) > 256 {
			fmt.Println("string is too long, limit is 256 symbols")
			continue
		}

		out := line[:len(line)-1]
		ch <- out
	}
}

func main() {
	ch := make(chan string)
	defer close(ch)

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
	}

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
	err = conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
