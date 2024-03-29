package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code        string
	Price       uint
	StudentName string `gorm:"student_name"` //默认就是student_name
}

func main() {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "root"  //密码
	host := "localhost" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "test"    //数据库名
	timeout := "10s"    //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//// 自动创建表
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100, StudentName: "小明"})

	// Read
	var product Product
	re := db.First(&product, 7) // 根据整型主键查找
	fmt.Println("返回结果：", re)
	fmt.Println("查询结果：", product)
	by, err := json.Marshal(product)
	fmt.Println(string(by))
	var producs []Product
	//First 是查找一个
	db.Find(&producs, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	data, err := json.Marshal(producs)
	fmt.Println(string(data))

	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)
}
