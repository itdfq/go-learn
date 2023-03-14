package main

import "fmt"

func main() {
	var x interface{}
	//写法一
	PrinthType(x)
	var x1 int
	PrinthType(x1)

	PrinthType(test)

	println("------------------")
	switch3(0)

}
func test(i int) {
	fmt.Println("打印：", i)
}
func PrinthType(x interface{}) {
	switch i := x.(type) { //初始化语句
	case nil:
		fmt.Printf("x 的类型为：%T\r\n", i)
	case int:
		fmt.Printf("x 的类型为int\n")
	case float64:
		fmt.Printf("x 是float64\n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 类型\n")
	case bool, string:
		fmt.Printf("x 是bool或者string型\n")
	default:
		fmt.Printf("x 是未知类型\n")
	}
}

func switch2(i int) {
	switch i {
	case 0:
	case 1:
		fmt.Printf("打印0或1")
	case 2:
		fmt.Printf("打印2")
	default:
		fmt.Printf("def")
	}
}

func switch3(i int) {
	switch i {
	case 0:
		println("fallthrough")
		fallthrough
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		println("继续执行")
	case 2:
		println("2")
	default:
		fmt.Printf("def")
	}
}
func switch4(i int) {
	switch {
	case i > 0 && i < 10:
		println("i>0 and i<10")
	case i >= 10 && i < 50:
		println("i>=10 and i<10")
	default:
		println("i>=50")
	}

}
