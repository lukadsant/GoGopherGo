package main

import (
	"fmt"
	"storesystem/store"
)

func main() {

	acc := store.Account{FirstName: "John", LastName: "Doe"}
	emp := store.Employee{Account: acc, Credits: 100.0}

	fmt.Println("Employee name", emp)

	emp.ChangeName("Jane", "Doe")
	fmt.Println("employer name after changename:", emp)

	emp.AddCredits(50.0)
	fmt.Println("credits after addcredits:", emp.CheckCredits())

	emp.RemoveCredits(30.0)
	fmt.Println("credits after removecredits:", emp.CheckCredits())

}
