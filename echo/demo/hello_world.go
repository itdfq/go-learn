package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http"
	"os"
)

type (
	User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

func main() {
	e := echo.New()
	//GET请求
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello,World")
	})

	//GET请求获取路径参数
	e.GET("/users/:id", getUser)

	//GET请求获取请求参数
	e.GET("/users", show)

	//GET 获取表单参数
	e.GET("/users", getForm)

	//POST方法
	e.POST("/post", func(c echo.Context) error {
		return c.String(http.StatusOK, "post test")
	})

	//POST方法上传文件测试
	e.POST("/post", getFile)

	//POST请求json
	e.POST("/getJson", Person)

	//设置静态字段目录
	//注意go语言中的上一层目录为 ./
	e.Static("/static", "././")

	//e.Start(":11323") 监听11323端口
	e.Logger.Fatal(e.Start(":11323"))

}
func getUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	//使用 Bind也可以直接获取form表单中的实体类型
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	fmt.Println("bind参数：", user)
	fmt.Println("接收到的参数：", id)
	return c.String(http.StatusOK, "接受到的参数为："+id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func Person(c echo.Context) error {
	var user []User

	result, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
		return err
	}

	err = json.Unmarshal(result, &user)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return err
	}
	headerName := c.Request().Header.Get("name")
	fmt.Println("请求头，name=", headerName)
	fmt.Println("user:", user) // user:[{1 xiaoming 13} {2 xiaohong 19}]
	//re, err := json.Marshal(user)
	//if err != nil {
	//	return nil
	//}
	return c.JSON(http.StatusOK, user)
}

func getForm(c echo.Context) error {
	//获取name 和 email 的值
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)

}

//form 表单 获取文件
func getFile(c echo.Context) error {
	//get name
	name := c.FormValue("name")
	//get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	//获取资源
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	//当函数返回结果时，执行，多个defer 逆序执行
	defer src.Close()
	filename := avatar.Filename
	size := avatar.Size
	fmt.Println("表单参数name:", name)
	fmt.Println("上传文件名字为：", filename)
	fmt.Println("上传文件的大小:", size/1024, "KB")

	//创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//copy
	if _, err = io.Copy(file, src); err != nil {
		return err
	}
	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")

}
