package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-learn/go/mysql/common"
)

func main() {
	var Db = common.Db
	r, err := Db.Exec("insert into person(username,sex,email) values(?,?,?)", "stu001", "man", "test@qq.com")
	common.HandleErr(err)
	id, err := r.LastInsertId()
	common.HandleErr(err)
	fmt.Println("insert success,id:", id)

	var person []common.Person
	err = Db.Select(&person, "select * from person where user_id = ?", 2)
	common.HandleErr(err)
	fmt.Println("select success:", person)

	res, err := Db.Exec("update person set username=? where user_id = ?", "韩信", 2)
	common.HandleErr(err)
	row, err := res.RowsAffected()
	common.HandleErr(err)
	fmt.Println("update success:", row)

	res, err = Db.Exec("delete from person where user_id = ?", 4)
	common.HandleErr(err)
	row, err = res.RowsAffected()
	common.HandleErr(err)
	fmt.Println("delete success:", row)

}
