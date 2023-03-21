package main

import (
	"fmt"
	"go-learn/go/socket/tcp"
	"net"
)

func main() {
	server()

}
func server() {
	lisen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	for {
		conn, err := lisen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}

		fmt.Println("客户端已连接")
		go tcp.Process(conn)
	}

}
