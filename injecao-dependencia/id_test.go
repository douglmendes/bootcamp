package main

import (
	"bytes"
	"testing"
)

func TestCumprimentar(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimentar(&buffer, "Chris")

	resultado := buffer.String()
	esperado := "Olá, Chris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
