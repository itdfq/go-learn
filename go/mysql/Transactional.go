package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-learn/go/mysql/common"
	_ "go-learn/go/mysql/common"
)

//
//import (“github.com/jmoiron/sqlx")
//   2)  Db.Begin()        开始事务
//   3)  Db.Commit()        提交事务
//   4)  Db.Rollback()     回滚事务
var Db = common.Db

//func init() {
//	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//	Db = database
//}

func main() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}

	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}
