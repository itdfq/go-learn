package main

import "fmt"

func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code = %d & endDate= %s"
	var targetUrl = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(targetUrl)

	fmt.Println("-----------------------------")

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}

	}

}
