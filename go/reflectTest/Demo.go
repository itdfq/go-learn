package reflectTest

import "fmt"

type Dog struct {
}

func (t Dog) T1(nums []int) {
	fmt.Println(nums)
	fmt.Println("t1")
}

func (t Dog) T2() {
	fmt.Println("t2")
}
