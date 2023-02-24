package main

import (
	"fmt"
	"sort"
)

//对于结构体来说，不是必须创建继承三个方法，才能实现排序
func main() {
	// 比较简单的例子，可以不用实现三个方法进行排序
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age }) // 按年龄升序排序
	fmt.Println("Sort by age:", people)

	sort.SliceStable(people, func(i, j int) bool { return people[i].Age > people[j].Age }) // 按年龄降序排序
	fmt.Println("Sort by age:", people)

	fmt.Println("Sorted:", sort.SliceIsSorted(people, func(i, j int) bool { return people[i].Age > people[j].Age })) //判断是否降序有序

	//查找
	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x := 21
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })           // 升序排序
	index := sort.Search(len(a), func(i int) bool { return a[i] >= x }) // 查找元素

	if index < len(a) && a[index] == x {
		fmt.Printf("found %d at index %d in %v\n", x, index, a)
	} else {
		fmt.Printf("%d not found in %v,index:%d\n", x, a, index)
	}

}
