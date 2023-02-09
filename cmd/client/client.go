package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Type message:")
		msg, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, msg+"\n")

		resp, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Response message:", resp)
	}
}
