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
func subtraction(s []int, c chan int) {
	sum := 100
	for _, v := range s {
		sum -= v

	}
	c <- sum
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	//默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
	c := make(chan int) //声明通道
	//执行 0 1 2
	fmt.Println("len(s)/2=", len(s)/2, "s[:len(s)/2]=", s[:len(s)/2])
	//i 取值：0，1，2
	go sum(s[:len(s)/2], c)
	//执行 2-最后 i取值 3，4，5
	go sum(s[len(s)/2:], c)
	go subtraction(s, c)

	//x, y := <-c, <-c
	//fmt.Println(x, y, x+y)
	x := <-c
	y := <-c
	//顺序是无法保证的
	println(x, y)
	println(<-c)
}
