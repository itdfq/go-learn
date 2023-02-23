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
