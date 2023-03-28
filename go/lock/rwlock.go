package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func main() {

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read1()

	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
func write() {
	rwlock.Lock() //加入写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock() //解写锁
	wg.Done()       //计数器-1
}

func read1() {
	rwlock.RLock() //加入读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() //解读锁
	wg.Done()        //计数器-1
}
