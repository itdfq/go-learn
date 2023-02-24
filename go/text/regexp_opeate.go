package main

import (
	"fmt"
	"log"
	"regexp"
)

//正则表达式
func main() {
	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}
	//匹配子字符串
	for _, word := range words {

		found, err := regexp.MatchString(".even", word)

		if err != nil {
			log.Fatal(err)
		}

		if found {

			fmt.Printf("%s matches\n", word)
		} else {

			fmt.Printf("%s does not match\n", word)
		}
	}
	fmt.Println("---------------解析字符串---------------")

	re, err := regexp.Compile(".even")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("解析的字符串：", re)
	for _, word := range words {
		found := re.MatchString(word)
		if found {
			fmt.Printf("%s matches\n", word)
		} else {

			fmt.Printf("%s does not match\n", word)
		}
	}

}
