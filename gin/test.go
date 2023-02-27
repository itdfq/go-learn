package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	//加入自定义的函数
	r.Use(StatCost())

	r.GET("/someJson", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hello world!"})

	})

	r.GET("/moreJson", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user""`
			Message string
			Age     int
		}
		msg.Message = "Hellp world"
		msg.Age = 18
		msg.Name = "小明"
		context.JSON(http.StatusOK, msg)

	})

	//获取QueryString参数
	r.GET("/user/search", func(context *gin.Context) {
		username := context.DefaultQuery("username", "小王子")
		fmt.Println("查询参数为：", username)
		address := context.Query("address")
		//输出json结果给调用方
		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})

	})

	//获取form参数
	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	//获取json参数
	r.POST("/json", func(c *gin.Context) {
		b, _ := c.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)

	})

	//获取path的参数
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	//http重定向
	r.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://baidu.com")
	})

	//路由重定向
	r.GET("/test1", func(c *gin.Context) {
		//指定重定向URL
		c.Request.URL.Path = "/test"
		r.HandleContext(c)
	})

	r.Run(":8087")
}

//定义中间件,计算耗时时间
func StatCost() gin.HandlerFunc {
	fmt.Println("执行了")
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("token", "假token")
		c.Next()
		cost := time.Since(start)
		log.Println(cost)
	}

}
