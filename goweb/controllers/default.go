package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goweb/models"
	"log"
	"path"
	"strings"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

type JsonController struct {
	beego.Controller
}
type XmlController struct {
	beego.Controller
}

type FileOptUploadController struct {
	beego.Controller
}

// 下载文件
type FileOptDownloadController struct {
	beego.Controller
}

//正则
type ZhengzeController struct {
	beego.Controller
}

type Result struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

//注解Controller
type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (c *CMSController) StaticBlock() {
	c.Ctx.WriteString("staticblock")
}

// @router /all/:key [get]
func (c *CMSController) AllBlock() {
	c.Ctx.WriteString("AllBlock")
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}

func (c *UserController) Get() {
	user := models.GetUserById(1)
	fmt.Println("查询的结果:", user)
	//模板数据数据
	c.Data["user"] = user
	c.Data["Website"] = "itdfq.com"
	c.Data["Email"] = "666.com"
	//使用新的实体模板
	//如果不指定，默认则会去 Controller/<方法名>.tpl 查找
	c.TplName = "user.tpl"

}

func (this *UserController) Post() {

	//直接输出字符串
	this.Ctx.WriteString("hello world!!!")
	//
}

//设置Controller 不需要屏蔽，Prepare（）方法在 Method 方法之前执行
func (this *UserController) Prepare() {
	this.EnableXSRF = false
}

// json格式输出
func (this *JsonController) Get() {
	mystruct := &Result{200, "", "测试数据", 0}
	this.Data["json"] = mystruct
	this.ServeJSON()
}

// xml格式输出
func (this *XmlController) Get() {
	xml := &Result{200, "", "xml格式数据", 0}
	this.Data["xml"] = xml
	this.ServeXML()

}

// 上传下载文件页面
func (c *FileOptUploadController) Get() {
	c.TplName = "fileopt.html"
}

// 上传文件
func (this *FileOptUploadController) Post() {
	//上传文件
	//image 对应的是html中name 的值   <input type="file" name="image" id="file_upload">
	f, h, _ := this.GetFile("image")
	//文件名称
	oldName := h.Filename
	fileName := h.Filename
	log.Print("当前上传文件为：", fileName)
	arr := strings.Split(fileName, ".")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	//关闭上传文件
	f.Close()
	//保存上传文件到指定位置
	this.SaveToFile("image", path.Join("static/uploadFile", oldName))
	log.Println("上传成功")
	//显示在本页面，不做跳转操作
	this.TplName = "fileopt.html"
}

func (this *FileOptDownloadController) Get() {
	//图片，txt,pdf，文件变成浏览了
	//this.Redirect("/static/0.jpg", 200)

	//直接下载
	this.Ctx.Output.Download("static/0.jpg", "图片1.jpg")

}
func (this *ZhengzeController) Get() {
	id := this.Ctx.Input.Param(":id")
	println("请求的id:", id)
	mystruct := &Result{200, "", "请求的id" + id, 0}
	this.Data["json"] = mystruct
	this.ServeJSON()

}

func (this *ZhengzeController) ListFood() {
	mystruct := &Result{200, "", "ListFood", 0}
	this.Data["json"] = mystruct
	this.ServeJSON()

}

func (this *ZhengzeController) GetFood() {
	///zhengze/getfood/2013/09/12
	params := this.Ctx.Input.Params()
	//"data": {
	//	"0": "2013",
	//		"1": "09",
	//		"2": "12",
	//		":splat": "2013/09/12"
	//}
	mystruct := &Result{200, "", params, 0}
	this.Data["json"] = mystruct
	this.ServeJSON()

}
