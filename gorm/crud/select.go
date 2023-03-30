package main

import (
	"encoding/json"
	"fmt"
	"go-learn/gorm/config"
	"go-learn/gorm/entry"
	"gorm.io/gorm/clause"
	"time"
)

func main() {
	db := config.ConfigMysql()
	var user entry.User
	//查询一条数据 默认主键升序
	db.First(&user) //SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println(user)

	//获取一条记录，没有制定排序字段
	db.Take(&user)
	fmt.Println(user) //select * from users limit 1

	//获取最后一条数据 (主键降序)
	db.Last(&user)
	fmt.Println(user) // select * from users order by id desc limit 1

	result := db.First(&user)
	rows := result.RowsAffected //影响的条数
	fmt.Println("受影响的行为：", rows)
	err := result.Error
	fmt.Println("错误信息:", err)
	//注意： First、Last 方法会根据主键查找到第一个、最后一个记录， 它仅在通过 struct 或提供 model 值进行查询时才起作用。 如果 model 类型没有定义主键，则按第一个字段排序

	fmt.Println("----------------------")
	result1 := map[string]interface{}{}
	db.Model(&user).First(&result1)
	data, err := json.Marshal(result1)
	fmt.Println(string(data))
	fmt.Println("---------------查询条件----------------")
	//主键查询
	//db.First(&user, 10)
	//db.First(&user, "10")
	var users []entry.User
	db.First(&users, []int{1, 2, 3}) //select * from user where id in (1,2,3)
	fmt.Println(users)

	fmt.Println("--------------基本条件查询------------------------")
	db.Where("name = ?", "jinzhu").Find(&users)

	db.Where("name in ?", []string{"jinzhu", "2"}).Find(&users)

	db.Where("name like ?", "%2%").Find(&users)

	db.Where("name=? and age>= ?", "jinzhu", 2).Find(&users)

	var lastweek, tody time.Time
	db.Where("created_at between ? and ?", lastweek, tody).Find(&users)

	fmt.Println("struct & Map 条件查询")
	db.Where(&entry.User{Name: "jinzhu", Age: 20}).Find(&users)

	db.Where(map[string]interface{}{"name": "jinzhu", "age": 10}).Find(&users)

	db.Where([]int64{10, 11, 12}).Find(&users)

	fmt.Println("----------------排序---------------")
	db.Order("age desc,name").Find(&users)
	db.Order("age desc").Order("name").Find(&users)
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&users) //SELECT * FROM users ORDER BY FIELD(id,1,2,3)

	fmt.Println("----------------分页----------------------")
	db.Limit(3).Find(&users)
	var users2 []entry.User
	//通过-1 消除limit条件
	db.Limit(10).Find(&users).Limit(-1).Find(&users2)
	//select * from users limit 10; (users)
	//select * from users; (users2)
	db.Limit(10).Offset(5).Find(&users)

	fmt.Println("--------------------group & Having---------------------")
	db.Model(&entry.User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name`

	db.Model(&entry.User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

	db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)")

	db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100)

}
