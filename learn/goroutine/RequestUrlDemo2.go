package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	ch := make(chan string) //声明一个通道
	for _, url := range apis {
		go reuqest(url, ch)
	}

	//time.Sleep(3 * time.Second)
	//for c := range ch {
	//	fmt.Printf("请求相应结果：%s", c)
	//}

	//需要注意
	for i := 0; i < len(apis); i++ {
		log.Printf("相应结果:%s\n", <-ch)
	}

	hostTime := time.Since(start)

	log.Println("花费时间为：", hostTime)
}
func reuqest(url string, c chan string) error {
	if "" == url {
		return errors.New("请求地址不能为空")
	}
	res, err := http.Get(url)
	if err != nil {
		log.Printf("请求url:%s,调用出错,原因：%v\n", url, err)
		c <- err.Error()
		return errors.New("请求出错")
	}
	bytes, _ := json.Marshal(res)
	log.Printf("请求地址为:%s->调用成功：%s\n", url, string(bytes))
	c <- url + "调用成功"
	return nil
}
