package main

import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"go-learn/reptile/reptile_common"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://www.baidu.com/")
	//reptile_common.HandleError(err)
	//body, err := ioutil.ReadAll(resp.Body)
	//reptile_common.HandleError(err)
	//fmt.Println(string(body))

	//第一种方式，不使用并发
	url := "http://www.baidu.com/"
	//download(url)

	//第二种方法，使用并发
	queue := make(chan string)
	go func() {
		queue <- url
	}()
	for uri := range queue {
		downloadV2(uri, queue)

	}
}
func download(url string) {
	clinet := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//自定义的Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := clinet.Do(req)
	reptile_common.HandleError(err)
	//延迟关闭函数链接
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		fmt.Println("parse url", link)
	}
}

func downloadV2(url string, queue chan string) {
	clinet := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := clinet.Do(req)
	reptile_common.HandleError(err)

	//延迟关闭
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)

	for _, link := range links {
		fmt.Println("download url", link)
		go func() {
			queue <- link
		}()
	}

}
