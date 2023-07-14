package main

import (
	"fmt"
	"sort"
)

//让map可以顺序输出
func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	n := map[string]int{
		"abe":    0,
		"orange": 1,
		"absd":   2,
	}
	keys1 := make([]string, 0, len(n))
	for k := range n {
		keys1 = append(keys1, k)
	}

	//对切片进行排序
	sort.Strings(keys1)

	for _, k := range keys1 {
		fmt.Printf("key=%s,value= %d \n", k, n[k])
	}

	// 将 map 中的键存储到切片中
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 按照排序后的顺序遍历 map
	for _, k := range keys {
		fmt.Printf("key=%s, value=%d\n", k, m[k])
	}

}
