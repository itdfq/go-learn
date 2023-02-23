package main

import "fmt"

func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code = %d & endDate= %s"
	var targetUrl = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(targetUrl)

}
