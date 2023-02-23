package main

import (
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"net/http"
	_ "net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello,World")
	})
	//e.Start(":11323") 监听11323端口
	e.Logger.Fatal(e.Start(":11323"))
	
}
