package main

import "fmt"

func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		//defer t.Close()  打印的结果全是：c  closed
		defer Close(t) //这个打印的是有顺序的

		//结论： defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。
		//但是并没有说struct这里的this指针如何处理，通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待。
	}

	testDefer()
}

func testDefer() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y) //y 闭包被引用
	}(x) //被复制

	x += 10
	y += 100
	println("x = ", x, "y = ", y)
}

type Test struct {
	name string
}

// 这个是方法
func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

//这个是函数
func Close(t Test) {
	t.Close()
}
