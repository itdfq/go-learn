package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-learn/reptile/reptile_common"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	reptile_common.HandleError(err)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	reptile_common.HandleError(err)
	fmt.Println((doc.Html()))
	fmt.Println(doc.Find("title").Text())
	defer resp.Body.Close()
}
