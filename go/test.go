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

}
