package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 09:14:00", time.Local)
	fmt.Println(time.Now().Sub(t).Hours())

}
