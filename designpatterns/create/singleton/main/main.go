package main

import (
	"fmt"
	"go-learn/designpatterns/create/singleton"
	"reflect"
)

func main() {
	ins1 := singleton.GetInstance(100)
	ins2 := singleton.GetInstance(200)
	if ins1 == ins2 {
		print("两个对象引用地址相等")

	}
	fmt.Println("--------------------------------")
	print("用DeepEqual判断：", reflect.DeepEqual(ins1, ins2), " \n")

	fmt.Println("对象1的值：", ins1.GetValue())
	fmt.Println("对象2的值：", ins2.GetValue())

	//结果

	/**
	两个对象引用地址相等--------------------------------
	用DeepEqual判断：true
	对象1的值： 100
	对象2的值： 100

	由于 once.Do() 这个方法只会调用一次，因此只有第一次100赋值成功了
	*/
}
