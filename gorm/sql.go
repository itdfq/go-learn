package main

import (
	"fmt"
	"go-learn/gorm/config"
	"go-learn/gorm/entry"
	"gorm.io/gorm"
)

func main() {
	var user entry.User
	db := config.ConfigMysql()
	db.AutoMigrate(user)
	db.Raw("select * from user where id = ?", 1).Scan(&user)
	fmt.Println(user)

	//原生查询
	res := db.Exec("update test.user set `name` = ?  , age = ? where id = ?", "小明SSS", 189, 1)
	fmt.Println("更新受影响的行为：", res.RowsAffected)
	//db.Exec("drop table user")

	//Exec SQL 表达式
	db.Exec("update user set age = ? where name =?", gorm.Expr("age * ?+?", 1, 1), "小明SSS")

}
