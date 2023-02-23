package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是：", countryCapitalMap[country])
	}
	//删除元素
	delete(countryCapitalMap, "France")

	fmt.Println("删除地图后")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	m := make(map[string]student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
