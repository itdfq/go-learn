package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     "121.36.77.21:6379",
		Password: "dfq5019.", // no password set
		DB:       0,          // use default DB
	})
	//通过client.ping() 检查是否成功连接到redis服务器
	_, err := Rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Rdb)
	Rdb.Set("test", "123", 100*time.Second)

	s := Rdb.Get("test")
	fmt.Println("获取的值:", s)
	fmt.Println("test", s.Val())

}
