// experimento que utiliza arrays para gerar uma media
package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 90
	x[1] = 80
	x[2] = 70
	x[3] = 86
	x[4] = 77

	var total float64 = 0

	for i := 0; i < 5; i++ {
		total += x[i]
	}

	fmt.Println(total / 5)
	mainshort()
}

func mainshort() {
	x := [5]float64{
		90,
		80,
		70,
		86,
		//77,
	}

	var total float64 = 0

	for i := 0; i < 5; i++ {
		total += x[i]
	}

	fmt.Println(total / 5)
}
