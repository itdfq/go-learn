package main

import (
	"fmt"
	"go-learn/gorm/config"
	"go-learn/gorm/entry"
	"gorm.io/gorm"
)

func main() {
	db := config.ConfigMysql()
	var user entry.User
	db.First(&user)
	fmt.Println(user)
	fmt.Println("======开始进行更新======")

	//更新单个列
	db.Model(&entry.User{}).Where("id =?", 1).Update("name", "小明二代")
	db.First(&user)
	fmt.Println(user)
	println("-------更新多个列------")
	db.Model(&user).Where("id = ?", 1).Updates(map[string]interface{}{"name": "小明三代", "age": 188})
	fmt.Println(user)

	//批量更新
	result := db.Table("user").Where("id in ?", []int{10, 11}).Updates(map[string]interface{}{"name": "更新之后的名字"})
	count := result.RowsAffected //更新的记录数
	fmt.Println("更新的条数:", count)

	//使用sql进行更新
	db.Model(&user).Update("age", gorm.Expr("age * ?+?", 2, 10))

	//更新单个列，不更新追踪时间
	db.Model(&user).UpdateColumn("name", "hello")

	//更新多个列
	db.Model(&user).UpdateColumns(entry.User{Name: "hello", Age: 18})

	//更新选中的列
	db.Model(&user).Select("name", "age").UpdateColumns(entry.User{Name: "hello", Age: 10})

}
