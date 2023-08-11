package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpTest(s chan string) {
	url := "http://localhost:8051/lms/aliexpressFullmanage/product/init"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "Cookie=JSESSIONID=15f1623253-5311-4985-b9f7-c921a7d4f240; LOGIN_IDENTITY=61646d696e5f313233343536; _ga=GA1.1.384835298.1690766344; SESSION=30cbf87d-f2db-4c67-ab76-a63eeabfe9c0")
	req.Header.Add("userInfo", "496376f32818df18140a7a7239b900a687da3dab14d1b725e4d9a7e144982dc65818c4497f1fd5d4d4d02d68f0fe471688c7684d192b3e3481f56cc322591004c6cfe93dc02b4140e9b9959eb33e2def7651215ba9a4d5c35552ce7ac992a096f36f310c3c2cf86b6ad94ac577468299cbf6e3f463951c1bded73f44bc8a8288babd4776281f47ed1530a53233702803fcab10bbfa1a980d2012f6b22a80e5a7df")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	s <- string(body)

}
func main() {
	ch := make(chan string, 10)
	for i := 0; i < 10; i++ {
		go HttpTest(ch)
	}

	for s := range ch {
		fmt.Println("请求结果：" + s)
	}
}
