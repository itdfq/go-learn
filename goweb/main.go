package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "goweb/routers"
)

func main() {

	beego.Run()
}
func init() {
	//初始化数据库连接
	// 这里注册一个default默认数据库，数据库驱动是mysql.
	// 第三个参数是数据库dsn, 配置数据库的账号密码，数据库名等参数
	//  dsn参数说明：
	//      username    - mysql账号
	//      password    - mysql密码
	//      db_name     - 数据库名
	//      127.0.0.1:3306 - 数据库的地址和端口
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8")

}
