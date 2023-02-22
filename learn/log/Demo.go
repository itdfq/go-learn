package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.SetPrefix("duan： ") //设置日志前缀
	//log.Fatalf("youyewjs") 	//使用 Fatalf 可以结束程序，相当于os.Exit(1)

	//log.Panic("end...") //日志消息，会获得错误堆栈跟踪。 后面的不会执行了

	fmt.Printf("hey,you are dog!!!")

	//将日志记录在文件中
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Hello ,I'm end")

}
