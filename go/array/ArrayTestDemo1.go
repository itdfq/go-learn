package main

import "fmt"

func main() {
	var array = [5]float32{10, 2, 3, 4, 21}
	for i, x := range array {
		println(i, x)
	}
	var balance [10]int
	balance[0] = 1
	print(balance[0])

	str := [5]string{"1", "2", "3", "a", "b"}
	for _, va := range str {
		print(va)
	}

	println("--------------------------------------")
	//多维数组
	tmp := [][]int{{1, 2, 3}, {3, 4, 5}, {4, 5, 6}, {4, 5, 6}}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("tmp[%d][%d]=%d\n", i, j, tmp[i][j])
		}

	}
	result := getAverage(tmp, 4, 3)
	println("平均值：", result)

}
func getAverage(arr [][]int, x int, y int) float32 {
	sum := 0
	count := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			//fmt.Printf("tmp[%d][%d]=%d\n", i, j, arr[i][j])
			sum = sum + arr[i][j]
			count++
		}

	}
	return float32(sum / count)
}
