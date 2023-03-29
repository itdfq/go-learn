package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//并发爬虫思路
/**
1. 初始化数据管道
2. 爬虫写出：26个协程向管道中添加图片链接
3. 任务统计协程：检查26个任务是否都完成，完成关闭通道
4. 下载协程：
*/

var (
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	//监控协程
	chanTask chan string
	reImage  = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	//DownloadFile("http://i1.shaodiyejin.com/uploads/tu/201909/10242/e5794daf58_4.jpg", "1.jpg")
	//1. 初始化管道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 26)
	//爬虫协程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	waitGroup.Add(1)
	go checkOK()
	//下载协程
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownLoadIma()

	}
	waitGroup.Wait()
}
func DownLoadIma() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}

	}
	waitGroup.Done()

}

// 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}
func checkOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

//获取图片url,并加入管道
func getImgUrls(url string) {
	urls := getImgs(url)
	for _, url := range urls {
		chanImageUrls <- url
	}
	//标识当前协程已完成，没完成一个任务，写一条数据
	chanTask <- url
	waitGroup.Done()
}

// 获取当前页图片链接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImage)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

func GetPageStr(url string) string {
	//1.去网站拿数据
	resp, err := http.Get(url)
	HandleError(err, "Http.Get url")
	//关闭
	defer resp.Body.Close()
	//2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//字节转字符串
	pageStr := string(pageBytes)
	return pageStr
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 下载图片，传入的是图片叫什么
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭异常")
		}
	}(resp.Body)
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	filename = "D:\\test\\go爬虫\\图片\\" + filename
	// 写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}
