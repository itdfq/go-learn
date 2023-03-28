package main

import (
	"fmt"
	"time"
)

func t1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func t2(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "test2"
}
func main() {
	//两个管道
	o1 := make(chan string)
	o2 := make(chan string)
	go t1(o1)
	go t2(o2)
	select {
	case s1 := <-o1:
		fmt.Println("s1=", s1)
	case s2 := <-o2:
		fmt.Println("s2=", s2)
	}

	o3 := make(chan string, 10)
	go write(o3)
	for s := range o3 {
		fmt.Println("res：", s)
		time.Sleep(time.Second)
	}

}
func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}

}
