package main

import "fmt"

func main() {
	s1 := test_a(func() int {
		return 100
	}) //匿名当前函数 作为参数
	s2 := format(func(s string, x, y int) string {
		//Sprintf根据格式说明符格式化并返回结果字符串。
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 10, 29)
	println(s1)
	println(s2)
}
func test_a(fn func() int) int {
	return fn()
}

//定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
