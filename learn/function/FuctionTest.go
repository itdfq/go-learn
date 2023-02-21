package main

func main() {
	read()
	count := max(1, 2)
	println(count)
	x, y := swap("a", "b")
	println(x, y)
	asd := 5
	asd = change(asd)
	println(asd)

}
func read() {
	println("小明正在读书")

}

func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

/**
函数可以返回多个值
*/
func swap(x, y string) (string, string) {
	return y, x
}
func change(a int) int {
	return a + 10
}
