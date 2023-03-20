package main

import (
	"fmt"
)

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func main() {
	u := User{1, "xiaomi"}
	u.Test()
	mV := u.Test
	mV() //隐式传递
	me := (*User).Test
	me(&u) //显式传递
}
