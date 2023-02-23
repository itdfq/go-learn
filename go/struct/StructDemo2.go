package main

import (
	"fmt"
	"log"
)

type Player struct {
	userid   int
	username string
	password string
}

//总结：方法 使用指针或者对象都可以调用成功
//     函数 必须得看方法参数
func main() {
	//第1种方式，先声明对象，再初始化
	var player1 Player
	player1.userid = 1
	player1.username = "lina1"
	player1.password = "123456"

	//第2种方式，声明同时初始化
	player2 := Player{2, "lina2", "123456"}

	//第3种方式，通过 field:value 形式初始化，该方式可以灵活初始化字段的顺序
	player3 := Player{username: "lina3", password: "123456", userid: 3}

	//上面三种初始化方式都是生产对象的，相应如果想初始化得到对象指针的三种方法如下：
	//第1种方式，使用 new 关键字
	player4 := new(Player)
	player4.userid = 4
	player4.username = "lina4"
	player4.password = "123456"

	//第2种方式，声明同时初始化
	player5 := &Player{5, "lina2", "123456"}

	//第3种方式，通过 field:value 形式初始化，该方式可以灵活初始化字段的顺序
	player6 := &Player{username: "lina3", password: "123456", userid: 6}
	fmt.Println(player1, player2, player3, player4, player5, player6)
	fmt.Println("--------------------------")
	//对象
	print_obj(player2)
	//print_ptr(player2)        //无法调用，编译出错
	player2.m_print_obj()
	player2.m_print_ptr()

	//指针
	//print_obj(player6)        //无法调用，编译出错
	print_ptr(player6)
	player6.m_print_obj()
	player6.m_print_ptr()
}

//传入 Player 对象参数  函数
func print_obj(player Player) {
	//player.username = "new"　　//修改并不会影响传入的对象本身
	log.Println("userid:", player.userid)
}

//传入 Player 对象指针参数 函数
func print_ptr(player *Player) {
	player.username = "new"
	log.Println("userid:", player.userid)
}

//接收者为 Player 对象的方法，方法接收者的变量，按照 GO 语言的习惯一般不用 this/self ，而是使用接收者类型的第一个小写字母，可以看标准库中的代码风格。
func (p Player) m_print_obj() {
	//p.username = "new"　　//修改并不会影响传入的对象本身
	log.Println("self userid:", p.userid)
}

//接收者为 Player 对象指针的方法
func (p *Player) m_print_ptr() {
	p.username = "new"
	log.Println("self userid:", p.userid)
}
