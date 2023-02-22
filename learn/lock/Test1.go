package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	//这个相当于Future
	var wg sync.WaitGroup
	//互互斥锁
	var mu sync.Mutex
	//十个协程数量
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			//1万叠加
			for j := 0; j < 10000; j++ {
				//加入互斥锁
				mu.Lock()
				count++
				//解锁
				mu.Unlock()
				//Mutex在大量并发的情况下，会造成锁等待，对性能的影响比较大
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
