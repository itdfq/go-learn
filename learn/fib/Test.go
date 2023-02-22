package main

import "fmt"

func main() {
	// 需求： 从控制台接收用户信息【姓名, 年龄, 薪水, 是否拿到offer】
	var name string
	var age byte
	var salary float32
	var isOffer bool

	fmt.Println("请输入姓名:")
	// 当程序执行到 fmt.Scanl(&name), 程序会停止这里, 等待用户输入, 并回车
	fmt.Scanln(&name)

	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水:")
	fmt.Scanln(&salary)

	fmt.Println("请输入是否拿到Offer:")
	fmt.Scanln(&isOffer)

	fmt.Printf("名字是 %v\n 年龄是 %v\n 薪水是 %v\n 是否拿到offer %v\n ", name, age, salary, isOffer)

	println("-----------------------------")

}
