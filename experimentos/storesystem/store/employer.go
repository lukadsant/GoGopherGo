package store

type Employee struct {
	Account
	Credits float64
}

func (e *Employee) AddCredits(amount float64) {
	e.Credits += amount
}

func (e *Employee) RemoveCredits(amount float64) {
	e.Credits -= amount
}

func (e Employee) CheckCredits() float64 {
	return e.Credits
}
