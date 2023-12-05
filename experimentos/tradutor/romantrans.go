package main

import (
	"fmt"
	"strings"
)

func romanToNumber(roman string) (int, error) {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0

	roman = strings.ToUpper(roman)

	for i := 0; i < len(roman); i++ {
		currentValue, exists := romanMap[string(roman[i])]
		if !exists {
			return 0, fmt.Errorf("caractere invalido %s", string(roman[i]))
		}

		if i+1 < len(roman) {
			nextValue, exists := romanMap[string(roman[i+1])]
			if !exists {
				return 0, fmt.Errorf("caractere invalido %s", string(roman[i+1]))
			}

			if nextValue > currentValue {
				result += nextValue - currentValue
				i++
				continue
			}
		}

		result += currentValue
	}

	return result, nil
}

func main() {

	var input string
	fmt.Print("digite uma roman numeral: ")
	fmt.Scan(&input)
	fmt.Println(romanToNumber(input))

}
