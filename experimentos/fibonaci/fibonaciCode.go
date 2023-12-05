package main

import (
	"fmt"
)

func fibonacciSeq(n int) []int {
	if n < 2 {
		return nil
	}

	fibonacci := []int{1, 1}

	for i := 2; i < n; i++ {
		nextTerm := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, nextTerm)
	}

	return fibonacci

}

func main() {
	var num int
	fmt.Print("qual sequencia de fibonacci você quer?")
	fmt.Scan(&num)
	fmt.Println("essa é a sequencia:", fibonacciSeq(num))
}
