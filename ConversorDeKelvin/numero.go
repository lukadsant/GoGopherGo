package main

import "fmt"

const tempK float64 = 400

func main() {

	//Convertendo temperatura de Kelvin para Celsius
	tempC := tempK - 273

	//imprimindo o resultado formatado
	fmt.Printf("O valor de ponto de ebulição de %.2fK é de %.2f°C\n", tempK, tempC)
}
