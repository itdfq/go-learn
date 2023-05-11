package main

import (
	"fmt"
	"time"
)

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

	println("--------------------------------")

	var c1, c2, c3 chan int
	var i1, i2 int
	go push(c1, 10)
	time.Sleep(2000)
	select {

	case i1 = <-c1:
		fmt.Println("received ", i1, "form c1")
	case c2 <- i2:
		fmt.Println("send ", i2, "to c2")
	case i3, ok := (<-c3):
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

}

func push(ch chan int, value int) {
	ch <- value

}
