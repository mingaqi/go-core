package main

import (
	"fmt"
	"go-core/_9_socket/_2nianbao/protocol"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		b, _ := protocol.Encode(msg)
		conn.Write(b)
	}
}
