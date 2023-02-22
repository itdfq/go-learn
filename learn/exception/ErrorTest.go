package main

import (
	"errors"
	"fmt"
)

func main() {
	a := -5.12
	result, e := Sqrt(float32(a))
	fmt.Println(result, e)

}

// Sqrt 相当于java的抛出异常
func Sqrt(f float32) (float32, error) {
	if f < 0 {
		return 0, errors.New("参数不能小于0")
	}
	return f * f, nil
}
