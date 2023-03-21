package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8008", nil)
	time.Sleep(1000)

}
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	//回复
	w.Write([]byte("www.itdfq.com"))

}
