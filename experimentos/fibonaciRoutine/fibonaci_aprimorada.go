package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, result chan <- float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	result <- x
}

func main() {
	start := time.Now()
	results :=make(chan float64, 15)
	for i := 1; i < 15; i++ {
		go fib(float64(i),results)
	}

	for i :=1; i < 15; i++{
		n := <- results
		fmt.Printf("Fib(%v): %v\n",i,n)
	}

	close(results)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
