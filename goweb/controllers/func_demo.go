package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type DemoController struct {
	beego.Controller
}

//初始化
func (this *DemoController) Init(ct *context.Context, childName string, app interface{}) {

}

func (this *DemoController) Prepare() { //这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。
	this.EnableXSRF = false //关闭XSRF验证
}

func (this *DemoController) Get() {

}

func (this *DemoController) Post() {

}

func (this *DemoController) Delete() {

}

func (this *DemoController) Put() { //如果用户请求的 HTTP Method 是 PUT，那么就执行该函数

}

func (this *DemoController) Head() { //如果用户请求的 HTTP Method 是 HEAD，那么就执行该函数

}

func (this *DemoController) Patch() { //用户请求的HTTP Method是PATCH，那么就执行该函数

}

func (this *DemoController) Options() { //用户请求的HTTP Method是OPTIONS，那么就执行该函数

}

func (this *DemoController) Finish() { //行完相应的 HTTP Method 方法之后执行的，默认是空

}

func (this *DemoController) Render() error { //这个函数主要用来实现渲染模板，如果 beego.AutoRender 为 true 的情况下才会执行。

	return nil
}
