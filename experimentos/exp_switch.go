package main

import "fmt"

func main() {

	i := 1

	for i <= 10 {

		switch i {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("ichi")
		case 2:
			fmt.Println("ni")
		case 3:
			fmt.Println("san")
		case 4:
			fmt.Println("go")
		default:
			fmt.Println("número desconhecido!")

		}
		i = i + 1
	}
}
