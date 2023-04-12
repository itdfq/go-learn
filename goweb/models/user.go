package models

import (
	"github.com/astaxie/beego/orm"
)

/*
*
CREATE TABLE `users` (

	`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
	`username` varchar(30) NOT NULL COMMENT '账号',
	`password` varchar(100) NOT NULL COMMENT '密码',
	 PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/
type User struct {
	Id       int
	Username string
	Password string
}

// 定义User绑定那个表
func (u *User) TableName() string {
	return "users"
}

// 初始化函数
func init() {
	//向orm注册user模型
	orm.RegisterModel(&User{})
}
func GetUserById(id int) *User {
	if id == 0 {
		return nil
	}
	//创建orm对象
	o := orm.NewOrm()
	//初始化一个User模型对象
	user := User{}
	user.Id = id
	//等价sql： select * from user where id = ?
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		return nil
	} else if err == orm.ErrMissPK {
		println("找不到主键")
		return nil
	}
	return &user
}
