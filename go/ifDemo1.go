package main

import "fmt"

func getStatus() int {
	return 123
}
func main() {
	if zt := getStatus(); zt != 0 {
		fmt.Println(zt)

	}
	nums := []int{2, 7, 11, 15}
	target := 9
	str := twoSum(nums, target)
	for k, v := range str {
		fmt.Println(k, "k->v", v)
	}
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			fmt.Println(p, ok)
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
