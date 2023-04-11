package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("yyy.txt", []byte("www.itdfq.com"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
