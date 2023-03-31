package main

import (
	"fmt"
	"go-learn/gorm/config"
	"go-learn/gorm/entry"
)

func main() {
	/**
	如果您的模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！

	拥有软删除能力的模型调用 Delete 时，记录不会被从数据库中真正删除。但 GORM 会将 DeletedAt 置为当前时间， 并且你不能再通过正常的查询方法找到该记录。
	*/
	db := config.ConfigMysql()
	//阻止全局删除
	var user entry.User
	err := db.Delete(&user).Error // gorm.ErrMissingWhereClause
	if err != nil {
		fmt.Println(err.Error())
	}
	//根据id 进行删除
	user.Id = 11
	db.Delete(&user)
	re := db.First(&user)
	if re != nil {
		fmt.Println(re)
	}
	fmt.Println(user)
	fmt.Println("--------------------------------")
	//带有额外条件的删除
	db.Where("name = ?", "jinzhu").Delete(&entry.User{})
	fmt.Println(user)
	//根据主键索引删除
	db.Delete(&user, []int{1, 2, 3})
	//查询被软删除的记录
	db.Unscoped().Where("age= 20").Find(&user)
	fmt.Println(user)
	//永久删除
	db.Unscoped().Where("name =?", "jinzhu").Delete(&entry.User{})
	fmt.Println(user)
}
