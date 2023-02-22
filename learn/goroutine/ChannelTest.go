package main

import "fmt"

/**
声明一个通道 ch
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把值发送到通道
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) //声明通道
	//执行 0 1 2
	fmt.Println("len(s)/2=", len(s)/2, "s[:len(s)/2]=", s[:len(s)/2])
	go sum(s[:len(s)/2], c)
	//执行 2-最后
	go sum(s[len(s)/2:], c)

	//x, y := <-c, <-c
	//fmt.Println(x, y, x+y)
	x := <-c
	y := <-c
	println(x, y)
}
