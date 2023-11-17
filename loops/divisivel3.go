package main

import "fmt"

func main() {
	// Inicializando a variável de controle do loop
	i := 1

	// Loop que percorre números de 1 a 101
	for i <= 101 {
		// Verifica se o número é divisível por 3
		if i%3 == 0 {
			// Se for divisível, imprime o número seguido por " : 0"
			fmt.Println(i, " : 0")
		}

		// Incrementa a variável de controle do loop
		i++
	}
}
