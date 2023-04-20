package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goweb/controllers"
	_ "goweb/routers"
	"html/template"
	"log"
	"net/http"
)

func main() {

	beego.Run()
}

const COLON = ":"

func init() {
	//初始化数据库连接
	// 这里注册一个default默认数据库，数据库驱动是mysql.
	// 第三个参数是数据库dsn, 配置数据库的账号密码，数据库名等参数
	//  dsn参数说明：
	//      username    - mysql账号
	//      password    - mysql密码
	//      db_name     - 数据库名
	//      127.0.0.1:3306 - 数据库的地址和端口

	//通过这种方式获取配置信息
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPasswd := beego.AppConfig.String("mysqlPasswd")
	mysqlUrls := beego.AppConfig.String("mysqlUrls")
	dbName := beego.AppConfig.String("mysqlDB")
	mysqlProt := beego.AppConfig.String("mysqlPort")
	dataSource := mysqlUser + COLON + mysqlPasswd + "@tcp(" + mysqlUrls + COLON + mysqlProt + ")/" + dbName + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dataSource)

	log.Println("数据库初始化成功,当前连接的数据库为：", dbName)
	runmode := beego.AppConfig.String("runmode")
	log.Println("系统启动成功,当前环境为：{" + runmode + "}")

	//自定义错误处理
	beego.ErrorHandler("404", page_not_found)

	//定义错误处理的controller
	beego.ErrorController(&controllers.FuncController{})

}
func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
