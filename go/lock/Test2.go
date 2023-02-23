package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/**
	RWmutex 读写锁
	读写锁的读锁可以重入，在已经有读锁的情况下，可以任意加读锁。
	在读锁没有全部解锁的情况下，写操作会阻塞直到所有读锁解锁。
	写锁定的情况下，其他协程的读写都会被阻塞，直到写锁解锁。

	可以再看看这篇文章： https://www.cnblogs.com/chenqionghe/p/13919427.html
	*/
	var m sync.RWMutex
	go read(&m, 1)
	go read(&m, 2)
	go read(&m, 3)

	time.Sleep(2 * time.Second)
}
func read(m *sync.RWMutex, i int) {
	fmt.Println(i, "reader start")
	m.RLock()
	fmt.Println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	fmt.Println(i, "reader over")
}
