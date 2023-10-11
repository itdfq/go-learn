package main

import (
	"fmt"
	"strconv"
)

//go的数据类型
func main() {
	// 布尔类型
	var b bool = true
	fmt.Println(b)

	//数字类型 float32 float64
	var a int = 123
	var a1 float32 = 12.34
	fmt.Println(a)
	fmt.Println(a1)
	//声明多个变量
	var x, y int
	x = 123
	y = 456
	fmt.Print(x + y)

	//变量赋值
	var x1 *int
	var x2 []int
	var x3 map[string]int = make(map[string]int)
	var x4 chan int
	var x5 func(string) int
	var x6 error
	fmt.Print("\n---------------------\n")
	for i := 1; i < 6; i++ {
		//int 转化成string 使用这种方式
		a := strconv.Itoa(i)
		fmt.Println("当前加入的元素的值为:", a)
		x3[a] = i
	}

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Printf("%v/n", x3)
	fmt.Println(x4)
	fmt.Println(x5)
	fmt.Println(x6)
	var i int
	var f float64
	var bx bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bx, s)
}
