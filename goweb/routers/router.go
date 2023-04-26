package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"goweb/controllers"
)

func init() {

	//固定路由
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})

	//测试json 返回数据
	beego.Router("/json", &controllers.JsonController{})

	beego.Router("xml", &controllers.XmlController{})

	beego.Router("/Home/FileOpt", &controllers.FileOptUploadController{})

	beego.Router("/Home/FileDown", &controllers.FileOptDownloadController{})

	//基本Get路由
	beego.Get("/get", func(context *context.Context) {
		context.Output.Body([]byte("Get基本路由"))
	})
	//基本Post路由
	beego.Post("/post", func(context *context.Context) {
		context.Output.Body([]byte("Post基本路由"))
	})
	//基本Put路由
	beego.Put("/put", func(context *context.Context) {
		context.Output.Body([]byte("基本路由"))
	})
	//基本Delete路由
	beego.Delete("/delete", func(context *context.Context) {
		context.Output.Body([]byte("基本路由"))
	})
	//基本Head路由
	beego.Head("/head", func(context *context.Context) {
		context.Output.Body([]byte("基本路由"))
	})
	//注册一个可以相应任何HTTP的路由
	beego.Any("/any", func(context *context.Context) {
		context.Output.Body([]byte("Any基本路由"))
	})

	/*********************正则路由*************************/
	// http://localhost:8999/api/123132
	beego.Router("/api/?:id", &controllers.ZhengzeController{})

	//自定义方法->匹配后面的方法
	beego.Router("/api/list", &controllers.ZhengzeController{}, "*:ListFood")
	beego.Router("/api/select", &controllers.ZhengzeController{}, "get,post:ListFood")
	beego.Router("/api/simple", &controllers.ZhengzeController{}, "get:GetFood;post:ListFood")

	/***********************自动匹配*************************/
	//会直接访问controller方法 全部小写生成路由;并且支持GET POST PUT DELETE
	beego.AutoRouter(&controllers.ZhengzeController{})
	//请求访问 /zhengze/listfood
	//zhengze/getfood

	/*********************注解路由****************************/
	//实验失败了，没有成功，具体原因待查看
	beego.Include(&controllers.CMSController{})

	/*******************获取数据*********************************/
	beego.Router("/getData", &controllers.DataController{})

	/****************过滤器**********************/
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")

		}
	}

	//beego.InsertFilter("/api/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/api/:id[0-9]+", beego.BeforeRouter, FilterUser)

	/**
	InsertFilter 函数的三个必填参数，一个可选参数

	pattern 路由规则，可以根据一定的规则进行路由，如果你全匹配可以用 *
	position 执行 Filter 的地方，五个固定参数如下，分别表示不同的执行过程
	BeforeStatic 静态地址之前
	BeforeRouter 寻找路由之前
	BeforeExec 找到路由之后，开始执行相应的 Controller 之前
	AfterExec 执行完 Controller 逻辑之后执行的过滤器
	FinishRouter 执行完逻辑之后执行的过滤器
	filter filter 函数 type FilterFunc func(*context.Context)
	*/

	//缓存
	beego.Router("/getCache", &controllers.CacheController{})
	beego.Router("/putCache", &controllers.CacheController{}, "post:PutCache")
}
