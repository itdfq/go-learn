package main

import "fmt"

func main() {
	// 需求： 从控制台接收用户信息【姓名, 年龄, 薪水, 是否拿到offer】
	var name string
	var age byte
	var salary float32
	var isOffer bool

	fmt.Println("请输入您的姓名, 年龄，薪水，是否拿到offer")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isOffer)
	fmt.Printf("姓名:%v\n 年龄:%v\n 薪水:%v\n 是否拿到offer:%v\n", name, age, salary, isOffer)
}
