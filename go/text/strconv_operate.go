package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "123"
	fmt.Println("原来的值:", s)
	/*字符串 转换 整形*/
	fmt.Println("字符串 转换 整形")
	//参数 base 代表字符串按照给定的进制进行解释。一般的，base 的取值为 2~36，
	//如果 base 的值为 0，则会根据字符串的前缀来确定 base 的值："0x" 表示 16 进制； "0" 表示 8 进制；否则就是 10 进制。

	//参数 bitSize 表示的是整数取值范围，或者说整数的具体类型。取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	fmt.Println("整形 转换 字符串  ")
	var a int64 = 456789
	result := strconv.FormatInt(a, 10)
	fmt.Println("转换结果：", result)

	fmt.Println("字符串 转换 布尔值  ")
	var bo string = "true"
	v, _ := strconv.ParseBool(bo)
	fmt.Println(v)

}
