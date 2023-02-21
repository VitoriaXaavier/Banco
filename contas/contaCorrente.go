package contas

import (
	"banco/cliente"
	"fmt"
)

type ContaCorrente struct {
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (cc *ContaCorrente) Deposita() {
	var deposita float64

	println("Digite o valor a ser depositado: ")
	fmt.Scan(&(deposita))
	if deposita > 0 {
		cc.saldo = deposita + cc.saldo
		fmt.Println("Valor após o deposito é de: ", cc.saldo)
	} else {
		fmt.Println("Valor de depósito menor que 0!")
	}
}

func (cc *ContaCorrente) Saque() {
	var saque float64

	println("Digite o valor para saque:")
	fmt.Scan(&(saque))

	if saque > 0 && saque <= cc.saldo {
		saque = cc.saldo - saque
		println("saldo restante: ", saque)

	} else {
		println("Valor indisponivel para saque!")
	}

}

func (cc *ContaCorrente) Transferir() float64 {
	var valorDaTransferencia float64
	var contaDestino *ContaCorrente

	if valorDaTransferencia < cc.saldo && valorDaTransferencia > 0 {
		println("Digite o valor a ser transferido")
		fmt.Scan(&(valorDaTransferencia))

		println("Qual a conta destino?")
		fmt.Scan(contaDestino)
		contaDestino = cliente.Cliente()

		cc.saldo -= valorDaTransferencia
		cc.Transferir()

		return valorDaTransferencia

	} else {
		println("Não foi possível realizar a transferencia")
		return cc.saldo
	}
}

