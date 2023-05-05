package main

import (
	"fmt"
	"go-learn/go/reflectTest"
	"reflect"
)

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func main() {
	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	fmt.Println("----------------------------------------")
	var dog reflectTest.Dog
	getType := reflect.TypeOf(dog)
	for i := 0; i < getType.NumMethod(); i++ {
		met := getType.Method(i)
		fmt.Printf("%s, %s, %s, %d\n", met.Type, met.Type.Kind(), met.Name, met.Index)
		fmt.Println(met.Func.Kind())

	}
	fmt.Println("------------------调用方法T2----------------------")
	met := getType.Method(1)
	met.Func.Call([]reflect.Value{reflect.ValueOf(dog)})
	fmt.Println("------------------调用方法T1----------------------")
	met1 := getType.Method(0)
	in := make([]reflect.Value, 1)
	//arg1 := 1
	//arg2 := 2
	//arg3 := 3

	arg := [3]int{1, 2, 3}
	in[0] = reflect.ValueOf(arg)
	//in[2] = reflect.ValueOf(arg2)
	//in[3] = reflect.ValueOf(arg3)
	met1.Func.Call(in)
}
