package main

import "fmt"

type Mover interface {
	move()
}
type Fox struct {
}

func (d Fox) move() {
	fmt.Println("狐狸会跑")
}

type People interface {
	Speak(string) string
}

// 接口可以嵌套，从而创造新的接口
type Animal interface {
	Sayer
	Mover
}

type Pig struct {
}

func (pig Pig) say() {
	fmt.Println("猪会say")
}
func (pig Pig) move() {
	fmt.Println("猪会move")
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
func main() {
	var x Mover
	var wangcai = Fox{}
	x = wangcai
	var fugui = &Fox{}
	x = fugui
	x.move()

	fmt.Println("--------------------")
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	fmt.Println("-------- 接口嵌套------------")
	var a Animal
	var p = Pig{}
	a = p
	a.say()
	a.move()

}
