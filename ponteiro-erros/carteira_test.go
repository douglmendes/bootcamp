package main

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	confirmarSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	confirmarErro := func(t *testing.T, resultado error, esperado error) {
		t.Helper()
		if resultado == nil {
			t.Error("Esperava um erro, mas nada ocorreu")
		}
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmarSaldo(t, carteira, Bitcoin(10))

	})
	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(5))
		confirmarSaldo(t, carteira, Bitcoin(15))
	})
	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmarSaldo(t, carteira, saldoInicial)
		confirmarErro(t, erro, ErroSaldoInsuficiente)
	})

}
