package main

import (
	"fmt"
	"go-learn/go/socket/tcp_sticky_package"
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
		b, err := tcp_sticky_package.Encode(msg)
		if err != nil {
			fmt.Println(err)
		}
		conn.Write(b)
	}
}
