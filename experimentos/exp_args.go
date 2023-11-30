//este experimento utiliza a função scanf
package main

import "fmt"

func main() {
	fmt.Print("digite um número: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

}
