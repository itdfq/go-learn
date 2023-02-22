package main

import "fmt"

func Send(ch chan string, message string) {
	ch <- message
}

func main() {
	size := 4
	ch := make(chan string, size)
	Send(ch, "one")
	Send(ch, "two")
	Send(ch, "three")
	Send(ch, "four")
	fmt.Println("All data sent to the channel ...")
	//当有缓存区的时候,所有的发送完毕,需要关闭,才能使用for循环获取具体的结果
	close(ch)
	for c := range ch {
		fmt.Println(c)

	}
	fmt.Println("Done!")

}
