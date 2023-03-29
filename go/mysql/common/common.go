package common

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

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

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("连接失败，原因：", err)
		return
	}
	Db = database
	fmt.Println("连接成功,", database)
}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
