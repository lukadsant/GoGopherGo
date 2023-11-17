package main

import "fmt"

func main() {
	i := 1
	for i <= 101 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("pinpan")
		} else if i%3 == 0 {
			fmt.Println("pin")
		} else if i%5 == 0 {
			fmt.Println("pan")
		}

		i++

	}
}
