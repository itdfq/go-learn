package main

import (
	"bufio"
	"fmt"
	"go-learn/go/socket/tcp_sticky_package"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//var buf [1024]byte
	for {
		//n, err := reader.Read(buf[:])
		message, err := tcp_sticky_package.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		//recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", message)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
