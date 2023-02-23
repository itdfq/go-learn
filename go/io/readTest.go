package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	// 从标准输入读取
	//data, err := ReadFrom(os.Stdin, 11)

	file, err := os.Open("info.log")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//获取文件的详情
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	size := fileInfo.Size()
	// 从普通文件读取，其中 file 是 os.File 的实例
	data, err := ReadFrom(file, size)

	// 从字符串读取
	//data, err = ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(string(data), err)

	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}

//Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误
//ReadFrom 函数将 io.Reader 作为参数，也就是说，ReadFrom 可以从任意的地方读取数据，只要来源实现了 io.Reader 接口。比如，我们可以从标准输入、文件、字符串等读取数据
func ReadFrom(reader io.Reader, num int64) ([]byte, error) {
	//创建指定的字节数组
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return nil, err
}
