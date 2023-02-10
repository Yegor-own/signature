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
		fmt.Print("Type message: ")
		msg, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, msg+"\n")

		resp, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println("Response message:", resp)
	}
}
