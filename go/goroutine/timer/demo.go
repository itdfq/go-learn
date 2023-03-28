package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	//timer基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	//验证timer 只能相应一次
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到")
	}



	//timer实现延时的功能
	//方式一
	time.Sleep(time.Second)
	//方式二
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("两秒时间到")
	//方式三
	<-time.After(2 * time.Second)
	fmt.Println("第二个两秒到了")


	//定时停止器
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		a := <-timer4.C
		fmt.Println("定时器执行了", a)
	}()

	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	}

	//重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

	*/

	//执行多个
	handlerMore()

}

func handlerMore() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				//停止
				ticker.Stop()
				fmt.Println("定时器执行停止")
			}
		}
	}()
	for {

	}
}
