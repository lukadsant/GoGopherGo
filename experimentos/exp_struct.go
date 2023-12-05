package main

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {

	var employee = Employee{
		Person: Person{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Address:   "123 Main St.",
		},
		ManagerID: 2,
	}

	employee.LastName = "doe"
	fmt.Println(employee)
}
