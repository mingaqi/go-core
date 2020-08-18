package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 1. 启动服务, 监听端口
	listenServe, err := net.Listen("_1tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("occurred err:", err)
		return
	}
	for {
		// 2. 建立链接
		conn, err := listenServe.Accept()
		if err != nil {
			fmt.Println("occurred err:", err)
			return
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	//buf := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	// 3. 读取数据
	for {
		msg, err := reader.ReadString('\n')
		//_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read occurred err:", err)
			return
		}
		fmt.Print(string(msg))
	}
}
