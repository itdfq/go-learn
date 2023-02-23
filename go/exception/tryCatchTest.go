package main

import "fmt"

func main() {
	errFun()
	println("执行完毕")
}
func errFun() {
	defer func() {
		//捕获异常
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	a := 1
	b := 0
	c := a / b
	fmt.Println(c)

}
