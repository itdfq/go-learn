package main

import "fmt"

func main() {
	var a float32

	a = 10000 / 6.0 / 30.0
	fmt.Println("平均一天需要赚：", a)
	fmt.Println("淀粉肠原价8元20根")
	fmt.Println("卖3元一根 5元两根")
	fmt.Println("20 根淀粉肠纯利润：", float32(20/2*5-8))
	fmt.Println("一根烤肠纯利润：", float32(20/2*5-8)/20)
	fmt.Println("----------------------------")
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示

}
