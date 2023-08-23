package main

func main() {
	a := transformMB(5368709120)
	println(a)

}

func transformMB(bit float32) float32 {
	if bit != 0 {
		return bit / 1024 / 1024
	} else {
		panic("输出参数有误!!!")
	}

}

//转换为GB
func transformGB(bit float32) float32 {
	if bit != 0 {
		return bit / 1024 / 1024 / 1024
	} else {
		panic("输出参数有误!!!")
	}

}
