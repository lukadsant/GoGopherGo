package main

import (
	"fmt"

	"github.com/lukadsant/calculator"
	"rsc.io/quote"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println("The sum of 3 and 5 is: ", total)
	fmt.Println("Version: ", calculator.Version)
	fmt.Println(quote.Hello())
}
