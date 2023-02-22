package main

import "fmt"

func main() {
	err := config("conf.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("执行完毕")
}

func config(conf string) error {
	if conf != "confi.ini" {
		return ErrNotFound
	} else {
		return nil
	}
}
