package face

import "fmt"

type student struct {
	name string
	age  int
}

//面试题
// 为什么v.name 都是博客？？？
func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	//for range 的 k,v 是固定的内存地址，只是循环的时候不断给它赋值，如果使用 &stu 那就是指向同一块内存地址。所以就会在最后一次赋值完成后，都是博客
	//解决办法在
	//for _, stu := range stus {
	//	stu := stu
	//	m[stu.name] = &stu
	//}
}
