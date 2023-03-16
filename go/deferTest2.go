package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func testLock() {
	lock.Lock()
	lock.Unlock()
}

func testDefer1() {
	lock.Lock()
	defer lock.Unlock()

}

func main() {
	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			testLock()
		}
		el := time.Since(t1)
		fmt.Println("test send: ", el)
	}()
	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			testDefer1()
		}
		el := time.Since(t1)
		fmt.Println("testDefer send :", el)
	}()
	// defer 执行在return之后 执行
	func() (i int) {
		i = 0
		defer func() {
			fmt.Println("执行了", i) //打印结果为：执行了 2
		}()
		return 2
	}()

	defer func() {
		//recover 是用来捕捉错误的
		//recover
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	testNil()

}

//名为 testNil 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。然而值得注意的是，run() 的声明是没有问题，因为在test函数运行完成后它才会被调用。
func testNil() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}
