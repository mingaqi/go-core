package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
loop:
	for {
		fmt.Print("say:")
		msg, _ := reader.ReadString('\n')

		if msg == "exit\n" {
			break loop
		}
		conn.Write([]byte(msg))

	}
	conn.Close()

}
