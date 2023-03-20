package main

import "fmt"

type Sayer interface {
	say()
}
type Cat struct {
}

type Dog struct {
}

func (cat Cat) say() {
	fmt.Println("猫-> 喵喵喵")
}

func (dog Dog) say() {
	fmt.Println("狗-> 汪汪汪")
}
