## 利用并发方法更快地计算斐波纳契数
使用以下程序按顺序计算斐波纳契数
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func fib(number float64) float64 {
    x, y := 1.0, 1.0
    for i := 0; i < int(number); i++ {
        x, y = y, x+y
    }

    r := rand.Intn(3)
    time.Sleep(time.Duration(r) * time.Second)

    return x
}

func main() {
    start := time.Now()

    for i := 1; i < 15; i++ {
        n := fib(float64(i))
    fmt.Printf("Fib(%v): %v\n", i, n)
    }

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
```

使用 channel 进行优化 
官方答案在：learn/fib/optimizeTest2.go