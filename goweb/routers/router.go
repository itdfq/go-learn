package routers

import (
	"github.com/astaxie/beego"
	"goweb/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})

	//测试json 返回数据
	beego.Router("/json", &controllers.JsonController{})

	beego.Router("xml", &controllers.XmlController{})

	beego.Router("/Home/FileOpt", &controllers.FileOptUploadController{})

	beego.Router("/Home/FileDown", &controllers.FileOptDownloadController{})
}
