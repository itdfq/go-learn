package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"

}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)

	for i := 0; i < 2; i++ {
		//多路复用，当有多个是，收到期中一个就会直接结束程序运行
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)

		}
	}
}
