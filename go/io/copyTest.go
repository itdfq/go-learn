package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开源文件
	srcFile, err := os.Open("info.log")
	if err != nil {
		panic(err)
	}
	//创建新文件
	dstFile, err := os.Create("info_copy.log")
	if err != nil {
		panic(err)
	}
	//缓冲读取
	buf := make([]byte, 1024)
	for {
		//从源文件读取
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//写出去
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()

}
