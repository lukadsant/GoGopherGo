package main

import (
	"errors"
	"fmt"
)

func checkNumber(num int) (string, error) {
	if num < 0 {
		return "", errors.New("number is negative")
	}

	return fmt.Sprintf("number %d é possitivo", num), nil
}

func main() {

	resultado, err := checkNumber(10)

	if err != nil {
		fmt.Println("Erro: ", err)

	} else {
		fmt.Println(resultado)
	}

	resultado, err = checkNumber(-10)
	if err != nil {
		fmt.Println("Erro: ", err)

	} else {
		fmt.Println(resultado)
	}

}
