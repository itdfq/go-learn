package main

import (
	"database/sql"
	"go-learn/gorm/config"
	"go-learn/gorm/entry"
	"time"
)

func main() {
	//获取数据库链接
	db := config.ConfigMysql()
	var email = "714845844@123.com"
	t1 := parse_timestr_to_timestamp("2023-03-30 14:15:00", 1)
	user := &entry.User{Name: "小明", Email: &email, Age: 17, Birthday: entry.JsonTime(t1), MemberNumber: sql.NullString{}}
	db.AutoMigrate(user)      //自动创建数据库表
	result := db.Create(user) //创建用户名
	println("创建数据库结果：", result)

	//批量添加
	//var users = []entry.User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	//db.Create(users)

	//或者使用 这个创建100条数据
	//db.CreateInBatches(users,100)

	//下面三种方式是跳过钩子函数的方法
	//db.Session(&gorm.Session{SkipHooks: true}).Create(&user)
	//db.Session(&gorm.Session{SkipHooks: true}).Create(&users)
	//db.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)

}
func parse_timestr_to_timestamp(time_str string, flag int) time.Time {
	var t time.Time
	loc, _ := time.LoadLocation("Local")
	if flag == 1 {
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, loc)
		t = t1
	} else if flag == 2 {
		t1, _ := time.ParseInLocation("2006-01-02 15:04", time_str, loc)
		t = t1
	} else if flag == 3 {
		t1, _ := time.ParseInLocation("2006-01-02", time_str, loc)
		t = t1
	} else if flag == 4 {
		t1, _ := time.ParseInLocation("2006.01.02", time_str, loc)
		t = t1
	} else {
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, loc)
		t = t1
	}
	return t
}
