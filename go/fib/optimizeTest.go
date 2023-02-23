package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib1(number float64, ch chan float64) float64 {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- x
	return x
}

func main() {
	var s string

	start := time.Now()
	ch1 := make(chan float64, 15)
	ch2 := make(chan string)
	for i := 1; i < 15; i++ {
		go fib1(float64(i), ch1)
	}
	go sendCancel(&s, ch2)
	for i := 1; i < 15; i++ {
		select {
		case a := <-ch1:
			fmt.Printf("Fib(%v): %v\n", i, a)
		case b := <-ch2:
			if b == "quit" {

				break
			} else {
				continue
			}
		}

	}
	t := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", t.Seconds())

}
func sendCancel(s *string, c chan string) {
	fmt.Scanln(s)
	c <- *s
}
