package tcp

import (
	"bufio"
	"fmt"
	"net"
)

//模仿TCP的服务端
func Process(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from Client failed,err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client 端发送过来的数据：", recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}
}
