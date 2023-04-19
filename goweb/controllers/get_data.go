package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type DataController struct {
	beego.Controller
}

func (this *DataController) Post() {
	//获取参数
	jsonInfo := this.GetString("jsonInfo")
	if jsonInfo == "" {
		this.Ctx.WriteString("jsoninfo is empty")
		this.StopRun() //提前终止进程
		return
	}
	//GetString(key string) string
	//GetStrings(key string) []string
	//GetInt(key string) (int64, error)
	//GetBool(key string) (bool, error)
	//GetFloat(key string) (float64, error)

	/************获取form表单结构数据*******************/
	u := user{}
	if err := this.ParseForm(&u); err != nil {
		//handler error
		fmt.Println("获取对象出错,msg:", err)
	}

	/*******************获取RequestBody内容***************************/
	//配置文件设置：copyrequestbody = true todo 待测试
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &u); err == nil {
		fmt.Println(u)
	} else {
		fmt.Println("获取requestBody失败,原因：", err)
	}

}

//设置表单结构
type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}
