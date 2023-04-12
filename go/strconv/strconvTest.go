package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转换为int
	s1 := "108"
	i1, err := strconv.Atoi(s1)
	handleError(err)
	fmt.Println("转换的值为：", i1)
	//int 转换为字符串
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Println("type:%T value:%v\n", s1, s2)
	//Parse 系列函数 它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	b1 := "T"
	b, err := strconv.ParseBool(b1)
	handleError(err)
	fmt.Println("type:%T value:%v\n", b1, b)

	//ParseInt base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	//
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	s3 := "1011001101"
	i3, err := strconv.ParseInt(s3, 2, 16)
	handleError(err)
	fmt.Println("type:%T value:%v\n", s3, i3)

	//ParseUnit

}
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
