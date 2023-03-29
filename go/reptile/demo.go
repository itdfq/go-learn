package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reQQEmail = `(\d+)@qq.com`

	// w代表大小写字母+数字+下划线
	reEmail = `\w+@\w+\.\w+`
	// s?有或者没有s
	// +代表出1次或多次
	//\s\S各种字符
	// +?代表贪婪模式
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	reVideo  = `https?://[^"]+?(\.((mp4)|(avi)|(rmvb)|(wmv)|(rm)))`
)

func main() {
	//getEmail("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	//getPhone("http://www.zhaohaowang.com/")
	//GetIdCard("http://list.chineseidcard.com/")
	//GetImg("https://juejin.cn/pins")
	//GetVideo("https://www.hi-oo.com/bf/?358817-5-2.html")
	GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
}

// 爬链接
func GetLink(url string) {
	pageStr := GetUrlResponseString(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[1])
	}
}

func getEmail(url string) {
	pageStr := GetUrlResponseString(url)
	//3. 过滤数据，获取邮箱
	re := regexp.MustCompile(reQQEmail)              //re := regexp.MustCompile(reStr)，传入正则表达式，得到正则表达式对象
	results := re.FindAllStringSubmatch(pageStr, -1) //ret := re.FindAllStringSubmatch(srcStr,-1)：用正则对象，获取页面页面，srcStr是页面内容，-1代表取全部
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}

}

func GetVideo(url string) {
	pageStr := GetUrlResponseString(url)
	println(pageStr)
	re := regexp.MustCompile(reVideo)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}

}

func GetImg(url string) {
	pageStr := GetUrlResponseString(url)
	//println(pageStr)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetIdCard(url string) {
	pageStr := GetUrlResponseString(url)
	println(pageStr)
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}
func getPhone(url string) {
	pageStr := GetUrlResponseString(url)
	fmt.Println("读取phone:", pageStr)
	//4. 获取手机号
	re1 := regexp.MustCompile(rePhone)
	phoneResults := re1.FindAllStringSubmatch(pageStr, -1)
	for _, phone := range phoneResults {
		fmt.Println("phone:", phone)

	}
}

func GetUrlResponseString(url string) string {
	//1.去网站拿数据
	resp, err := http.Get(url)
	HandlerError(err, "Http.Get url")
	//关闭
	defer resp.Body.Close()
	//2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandlerError(err, "ioutil.ReadAll")
	//字节转字符串
	pageStr := string(pageBytes)
	return pageStr
}
func HandlerError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
