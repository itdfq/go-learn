package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var one sync.Once

func main() {
	a := 0.01
	sum := 0.00
	for i := 0; i < 30; i++ {
		sum = sum + a
		a = a * 2
	}
	fmt.Printf("最后一天的数量为：%f\n", a)
	fmt.Printf("%f\n", sum)
	fmt.Println("---------------------------")
	wg.Add(1)
	go hello()
	fmt.Println("main groutine done!")
	wg.Wait()

	fmt.Println("测试sync.One 只会执行一次方法")
	syncTest()

}
func hello() {
	defer wg.Done()
	fmt.Println("hello Groutine!")
}
func syncTest() {
	for i := 1; i < 5; i++ {
		one.Do(func() {
			fmt.Println("执行do方法")
		})

	}

}
