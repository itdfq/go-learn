package main

import "fmt"

func main() {
	a := 5
	fmt.Printf("%b\n", a)
	//右移运算符
	a = a >> 2
	fmt.Printf("5 右移结果： %b\n", a)
	//左移运算符
	a = 5
	a = a << 2
	fmt.Printf("5 左移结果： %b\n", a)
	//小数点不动，数 动

	a = 5
	a = a << 3
	fmt.Printf("%b\n", a)

}
