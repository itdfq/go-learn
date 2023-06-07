package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// 例题: golang 实现RPC程序，实现求矩形面积和周长

type Params struct {
	Width, Height int
}
type Rect struct {
}

type Params1 struct {
	X, Y int
}

//RPC 服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	log.Println("收到求矩形面积的方法,当前获取的参数为:", p)
	*ret = p.Height * p.Width
	return nil
}

//周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	log.Println("收到求矩形周长的方法,当前获取的参数为:", p)
	*ret = (p.Width + p.Height) * 2
	return nil
}

func (r *Rect) Addition(p Params1, result *int) error {
	log.Println("收到求矩形面积的方法,当前获取的参数为:", p)
	*result = p.Y * p.Y
	return nil
}

//主函数
func main() {
	//注册服务
	rect := new(Rect)
	//注册一个 rect的服务
	rpc.Register(rect)
	// 服务处理绑定到http协议上
	rpc.HandleHTTP()
	//监听服务
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("服务启动成功")
}
