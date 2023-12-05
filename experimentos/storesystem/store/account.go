package store

import "fmt"

type Account struct {
	FirstName string
	LastName  string
}

func (a Account) String() string {
	return fmt.Sprintf("Account: %s %s", a.FirstName, a.LastName)
}

func (a *Account) ChangeName(newFirstName, newLastName string) {
	a.FirstName = newFirstName
	a.LastName = newLastName
}
