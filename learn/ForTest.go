package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	println(sum)

	count := 1
	for count < 10 {
		count = count + sum

	}
	println(count)

	//无限循环语句
	//for {
	//	println("循环")
	//}

	//尝试循环数组
	strings := []string{"A", "B", "C", "D"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

}
