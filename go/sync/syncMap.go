package main

import (
	"fmt"
	"strconv"
	"sync"
)

//go 语言内置的map 不是线程安全的
//var m = make(map[string]int)

//func get(key string) int {
//	return m[key]
//}
//func set(key string, value int) {
//	m[key] = value
//}
var m = sync.Map{} // sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法

func main() {
	//会报错
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			//set(key, n)
			m.Store(key, n)
			//fmt.Printf("k=%v,v=%v\n", key, get(key))
			value, _ := m.Load(key)

			fmt.Printf("k=%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
