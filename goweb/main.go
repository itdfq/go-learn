package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	"goweb/controllers"
	"goweb/models"
	_ "goweb/routers"
	"html/template"
	"net/http"
)

var log logs.BeeLogger

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

	log.Info("数据库初始化成功,当前连接的数据库为：", dbName)
	runmode := beego.AppConfig.String("runmode")
	log.Info("系统启动成功,当前环境为：{" + runmode + "}")

	//自定义错误处理
	beego.ErrorHandler("404", page_not_found)

	//定义错误处理的controller
	beego.ErrorController(&controllers.FuncController{})

	//开启sql查询debug
	orm.Debug = true

	//日志框架
	log := logs.NewLogger()
	//日志引擎设置
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	log.EnableFuncCallDepth(true) //输出文件名和文件号
	log.Debug("服务启动成功")
	log.SetLogFuncCallDepth(2)

	logs.Async()    //设置输出日志，
	logs.Async(1e3) //设置异步输出的缓冲池chan的大小

}

func crud() {
	o := orm.NewOrm()
	//insert
	user := models.User{Username: "123"}
	o.Insert(user)

	//update
	user.Username = "update-name"
	num, err := o.Update(&user)
	fmt.Println("修改受影响的行数:%d,报错:%v\n", num, err)

	//read one
	u := models.User{Id: user.Id}
	err = o.Read(&u)
	fmt.Println("readone：", u)

	//delete
	num, err = o.Delete(&u)

	//关联查询
	var posts []*models.Post
	qs := o.QueryTable("post")
	num, err = qs.Filter("User_Name", "slence").All(&posts)
	fmt.Println("关联查询的num:", num)

	//sql 查询
	var maps []orm.Params
	num, err = o.Raw("select * from user").Values(&maps)
	for _, term := range maps {
		fmt.Println(term["id"], ":", term["name"])
	}

	//事务处理
	o.Begin()
	//处理具体的业务逻辑
	user = models.User{Username: "username"}
	id, err := o.Insert(&user)
	if err != nil {
		o.Commit()
		fmt.Println("插入的id为：", id)

	} else {
		o.Rollback()
	}

}
func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

//初始化session
func init_session() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionID", //传入客户端存储Cookie的名字
		EnableSetCookie: true,
		Gclifetime:      3600, //
		Maxlifetime:     3600,
		Secure:          false, //是否开启HTTPS
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp", //配置设置，设置引擎设置
	}
	globalSessions, _ := session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

//模拟登录
func login(w http.ResponseWriter, r *http.Request) {
	/**
	函数介绍
	SessionStart 根据当前请求返回 session 对象
	SessionDestroy 销毁当前 session 对象
	SessionRegenerateId 重新生成 sessionID
	GetActiveSession 获取当前活跃的 session 用户
	SetHashFunc 设置 sessionID 生成的函数
	SetSecure 设置是否开启 cookie 的 Secure 设置
	*/
	sess, _ := beego.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w) //release the resource & save data to provider & return the data
	userName := sess.Get("username")
	fmt.Println("当前登录的用户名称为", userName)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		sess.Set("username", userName)
	}

}
