package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
CREATE TABLE `person` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(260) DEFAULT NULL,
    `sex` varchar(260) DEFAULT NULL,
    `email` varchar(260) DEFAULT NULL,
    PRIMARY KEY (`user_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE place (
    country varchar(200),
    city varchar(200),
    telcode int
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
*/

var Db *sqlx.DB

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}
type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func main() {
	r, err := Db.Exec("insert into person(username,sex,email) values(?,?,?)", "stu001", "man", "test@qq.com")
	HandleErr(err)
	id, err := r.LastInsertId()
	HandleErr(err)
	fmt.Println("insert success,id:", id)

	var person []Person
	err = Db.Select(&person, "select * from person where user_id = ?", 2)
	HandleErr(err)
	fmt.Println("select success:", person)

	res, err := Db.Exec("update person set username=? where user_id = ?", "韩信", 2)
	HandleErr(err)
	row, err := res.RowsAffected()
	HandleErr(err)
	fmt.Println("update success:", row)

	res, err = Db.Exec("delete from person where user_id = ?", 4)
	HandleErr(err)
	row, err = res.RowsAffected()
	HandleErr(err)
	fmt.Println("delete success:", row)

}
func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("连接失败，原因：", err)
		return
	}
	Db = database
	fmt.Println("连接成功,", database)
}
