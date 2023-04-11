package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//bufio包实现了带缓冲区的读写，是对文件读写的封装
//bufio缓冲写数据

func main() {
	//wr()
	//读取测试
	re()

}

func wr() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	//延迟关闭file
	defer file.Close()
	//获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("测试使用 bufio读取文件\n")
	}
	//刷新缓冲区，强制写出
	writer.Flush()

}

func re() {
	file, err := os.Open("info.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}
