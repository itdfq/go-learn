package main

import "fmt"

func main() {
	var x interface{}
	x = "itdfq.cn"
	//断言格式： x.(T)
	//    x：表示类型为interface{}的变量
	//    T：表示断言x可能是的类型。
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
