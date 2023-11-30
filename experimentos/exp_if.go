package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i, "Par")
		} else {
			fmt.Println(i, "Impar")
		}
		i = i + 1
	}
}
