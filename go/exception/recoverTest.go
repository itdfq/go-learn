package main

import "fmt"

func main() {
	//注意使用recover 需要使用defer 延迟函数，最后在执行
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//		//打印日志
	//		panic(err)
	//	}
	//}()
	//
	//var ch chan int = make(chan int, 10)
	//close(ch)
	//ch <- 1

	//testDeferException() //捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil    结果： defer panic

	testMoreException()
}

func testDeferException() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

//捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
func testMoreException() {
	defer func() {
		fmt.Println("第一", recover()) //有效
	}()

	defer recover() //无效

	defer fmt.Println("第二", recover()) //无效

	defer func() {
		func() {
			println("defer inner")
			recover() //无效
		}()
	}()
	panic("test panic")

}
