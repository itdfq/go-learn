package main

import (
	"fmt"
	"regexp"
)

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

	fmt.Println("=========================")
	s := "130821199103278829"
	re := regexp.MustCompile("[123456789]\\d{5}((19\\d{2})|(20[01]\\d))((0[1-9])|(1[012]))((0[1-9])|([12]\\d)|(3[01]))\\d{3}[\\dXx]")
	result := re.FindAllStringSubmatch(s, -1)
	for _, v := range result {
		fmt.Println(v)
	}

}
