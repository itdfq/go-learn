package main

import (
	"fmt"
	"time"
)

func main() {
	//goroutine 相当于开启一个线程去执行
	go say("world")
	say("hello")
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}
