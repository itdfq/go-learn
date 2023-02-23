package main

import "fmt"

func main() {
	ch := make(chan int, 2) //定义缓存区大小为2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	c := make(chan int, 10)
	go fi(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fi(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//使用完之后,必须得关闭，不然会一直阻塞，因为使用了range 循环遍历，如果不关闭，会一直阻塞
	close(c)
}
