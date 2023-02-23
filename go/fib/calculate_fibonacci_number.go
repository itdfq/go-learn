package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float32) float64 {
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
		n := fib(float32(float64(i)))
		fmt.Printf("Fib(%v): %v\n", i, n)
	}
	t := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", t.Seconds())
}
