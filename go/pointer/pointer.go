package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("a变量的地址:%x\n", &a)
	var ip *int
	ip = &a
	fmt.Printf("ip变量地址：%x\n", ip)

	//使用指针访问值
	fmt.Printf("ip实际的值：%d\n", *ip)

}
