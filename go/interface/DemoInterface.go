package main

import "fmt"

type Phone interface {
	call()
	count(x, y int) int
}
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you!")
}

func (nokiaPhone NokiaPhone) count(x, y int) int {
	fmt.Println("NokiaPhone当前计算的x:", x, " y:", y, " 计算结果=", x*y)
	return x * y
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iphone,I can call you!")
}
func (iphone IPhone) count(x, y int) int {
	fmt.Println("IPhone当前计算的x:", x, " y:", y, " 计算结果=", x+y)
	return x + y
}

//结构体
type User struct {
	Name  string
	Email string
}

//方法
func (u User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}
func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone.count(1, 2)

	phone = new(IPhone)
	phone.call()
	phone.count(1, 2)
	fmt.Println("--------------------")
	u1 := User{"name", "458444@qq.com"}
	u1.Notify()
}
