package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Chris")
	esperado := "Olá, Chris"

	if resultado != esperado {
		t.Errorf("Resultado '%s' é diferente do esperado '%s'", resultado, esperado)
	}
}
