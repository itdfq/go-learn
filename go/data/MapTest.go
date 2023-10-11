package main

import "fmt"

func main() {
	var siteMap map[string]string
	siteMap = make(map[string]string)
	siteMap["Google"] = "谷歌"
	siteMap["baidu"] = "百度"
	for site := range siteMap {
		println(site, "-->", siteMap[site])
	}
	fmt.Printf("当前元素的map为：%v", siteMap)
	//获取期中一个元素
	name, ok := siteMap["baidu"]
	if ok {
		fmt.Println("FaceBook的站点是", name)
		fmt.Println(name, ok)
	} else {
		fmt.Println("站点不存在")
		fmt.Println(name, ok)

	}

}
