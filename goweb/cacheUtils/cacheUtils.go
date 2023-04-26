package cacheUtils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"time"
)

var redisCache cache.Cache

func init() {
	//设置memory缓存
	//bm, err := cacheUtils.NewCache("memory", `{"interval":60}`)
	//设置file 配置 CachePath 表示缓存的文件目录，FileSuffix 表示文件后缀，DirectoryLevel 表示目录层级，EmbedExpiry 表示过期设置
	//bm, err := cacheUtils.NewCache("file", `{"CachePath":"./cacheUtils","FileSuffix":".cacheUtils","DirectoryLevel":"2","EmbedExpiry":"120"}`)

	//设置redis
	redisHost := beego.AppConfig.String("redis::address")
	dataBase := beego.AppConfig.String("redis::database")
	password := beego.AppConfig.String("redis::password")
	redisKey := beego.AppConfig.String("redis::key")
	config := fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%s","password":"%s"}`, redisKey, redisHost, dataBase, password)
	var err error
	redisCache, err = cache.NewCache("redis", config)
	if err != nil {
		panic(err)
	}
	var key = "astaxie"
	redisCache.Put(key, 1, 100*time.Second)
	value := redisCache.Get(key)
	fmt.Println(key, "的值为:", value)
	isExist := redisCache.IsExist(key)
	fmt.Println(key, "是否存在,result:", isExist)

	//删除key
	//bm.Delete(key)
}

func SetStr(key, value string, time time.Duration) (err error) {
	err = redisCache.Put(key, value, time)
	if err != nil {
		beego.Error("set key:", key, ",value:", value, err)
	}
	return
}

func GetStr(key string) (value string) {
	v := redisCache.Get(key)
	if v == nil {
		return ""
	}
	value = string(v.([]byte)) //这里的转换很重要，Get返回的是interface
	return
}
func DelKey(key string) (err error) {
	err = redisCache.Delete(key)
	return
}
