package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type Player struct {
	userid   int
	username string
	password string
}

func main() {
	//创建一个新的结构体
	fmt.Println(Books{"Go", "duan", "go", 1})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	book := Books{title: "Go 语言", author: "www.runoob.com"}
	// 忽略的字段为 0 或 空
	fmt.Println(book)
	fmt.Println(book.title)
	printBook(book)
	printBook1(&book)

	fmt.Println("-------------------------------")

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
}
func printBook(books Books) {
	fmt.Println(books)

}

//结构体指针
func printBook1(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
