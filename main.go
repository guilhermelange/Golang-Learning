package main

import (
	"Banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)
	PagarBoleto(&contaExemplo, 60)
	fmt.Println(contaExemplo.ObterSaldo())
}
