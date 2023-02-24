package main

import (
	"fmt"
	"sort"
)

func main() {
	x := 11
	s := []int{3, 6, 8, 11, 45} // 注意已经升序排序

	//该方法会使用“二分查找”算法来找出能使 f(x)(0<=x<n) 返回 ture 的最小值 i。 前提条件 : f(x)(0<=x<i) 均返回 false, f(x)(i<=x<n) 均返回 ture。 如果不存在 i 可以使 f(i) 返回 ture, 则返回 n。
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	fmt.Println("查询结果为：", pos)
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素 ", x)
	}
	GuessingGame()
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number >= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
