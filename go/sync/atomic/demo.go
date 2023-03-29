package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var l sync.Mutex
var wg sync.WaitGroup

func add() {
	x++
	wg.Done()
}

func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		//go add() //普通版add  不是线程安全的      95191   31338400
		//go mutexAdd() //枷锁版add  线程安全    100000  30286600
		go atomicAdd() //                      100000  29423100  性能比加锁好
	}
	wg.Wait()
	end := time.Now()
	println(x)
	println(end.Sub(start))
}
