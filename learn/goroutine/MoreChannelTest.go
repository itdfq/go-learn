package main

import (
	"fmt"
)

/**
多个channel使用举例，例如一个用于发送，一个用于接收
chan<- int // it's a channel to only Send data
<-chan int // it's a channel to only receive data
*/
func main() {
	ch := make(chan string, 1)
	send1(ch, "Hello World!")
	read1(ch)
}
func send1(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read1(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}
