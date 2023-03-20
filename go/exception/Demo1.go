package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func test01() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	index := 10
	a[index] = 10
	fmt.Println(a)
}
func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		panic("半径不能未负数")
	}
	return 3.14 * radius * radius
}
func test02() {
	getCircleArea(-5)
}
func test03() {

	//捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("执行到此")
}
func test04() {
	test03()
	fmt.Println("test04")
}
func getCircleArea2(radius float32) (area float32, err error) {
	if radius < 0 {
		// 构建个异常对象
		err = errors.New("半径不能为负2")
		return 0, err
	}
	area = 3.14 * radius * radius
	return area, nil
}

type pathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p pathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func Open(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return pathError{fileName, "read", time.Now().String(), err.Error()}
	}
	defer file.Close()
	return nil
}

func main() {
	test04()
	//方法自己抛出的异常
	area, err := getCircleArea2(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
	fmt.Println("--------------")
	err = Open("/test/asd")
	switch v := err.(type) {
	case pathError:
		fmt.Println("get Path Error,", v)
	default:
		fmt.Println("others,", v)
	}

}
