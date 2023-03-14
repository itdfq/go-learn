package main

import (
	"fmt"
	"time"
)

//select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
var resChan = make(chan int) //注意开启的是 无缓存的通道；main线程执行写入的时候，会一直阻塞等到去取出,后面的方法不会进行执行；解决办法，开启一个协程，进行通道写入

func main() {
	go fuzhi(3, resChan)
	test1()
}

func fuzhi(i int, c chan int) {
	c <- i
}
func test1() {
	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second * 3):
		fmt.Printf("超时3秒")
	}
}

func doData(data int) {

	println("执行的参数：", data)
}
