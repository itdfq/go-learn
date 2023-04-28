package main

import (
	"fmt"
)

//抽象工厂
//定义抽象工厂，有小米工厂和苹果工厂， 小米可以卖手机和手表；苹果工厂也可以卖手机和手表

type Phone interface {
	//手机可以用来打游戏
	playGame()
}
type Watch interface {
	//手表可以用来看时间
	see()
}

//定义抽象工厂
type AbstractFactory interface {
	productPhone() Phone
	productWatch() Watch
}

type MiPhone struct{}
type Iphone struct{}
type MiWatch struct{}
type IWathch struct{}

type MiFactory struct{}
type IFactory struct{}

func (m *MiPhone) playGame() {
	fmt.Println("小米手机打王者")
}
func (i *Iphone) playGame() {
	fmt.Println("苹果手机打原神")
}
func (m MiWatch) see() {
	fmt.Println("小米手表看时间")
}
func (i IWathch) see() {
	fmt.Println("苹果手表看日期")
}

func (f IFactory) productPhone() Phone {

	fmt.Println("苹果工厂生产苹果手机")
	return &Iphone{}
}

func (f IFactory) productWatch() Watch {
	fmt.Println("苹果工厂生成苹果手表")
	return &IWathch{}
}

func (f *MiFactory) productPhone() Phone {
	fmt.Println("小米工厂生成小米手机")
	return &MiPhone{}
}
func (f *MiFactory) productWatch() Watch {
	fmt.Println("小米工厂生成小米手表")
	return &MiWatch{}
}

func main() {
	var factory AbstractFactory
	var phone Phone
	var watch Watch
	//todo  这里暂时写死  后面考虑用反射获取值

	factory = new(MiFactory)
	phone = factory.productPhone()
	watch = factory.productWatch()
	phone.playGame()
	watch.see()
}
