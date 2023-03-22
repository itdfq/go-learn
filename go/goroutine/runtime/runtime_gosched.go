package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	//主线程
	for i := 0; i < 2; i++ {
		//Gosched产生了处理器，允许其他goroutine运行。它不会挂起当前的goroutine，因此执行会自动恢复
		runtime.Gosched() //切换时间片 加了这个world 才能打印出来
		fmt.Println("hello")
	}
}
