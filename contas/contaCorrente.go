package contas

import (
	"fmt"
	"home/vitoria/workspace/go/src/banco/cliente"

)
type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (cc *ContaCorrente) deposita() {
	var deposita float64

	println("Digite o valor a ser depositado: ")
	fmt.Scan(&(deposita))
	if deposita > 0 {
		cc.Saldo = deposita + cc.Saldo
		fmt.Println("Valor após o deposito é de: ", cc.Saldo)
	} else {
		fmt.Println("Valor de depósito menor que 0!")
	}
}

func (cc *ContaCorrente) saque() {
	var saque float64

	println("Digite o valor para saque:")
	fmt.Scan(&(saque))

	if saque > 0 && saque <= cc.Saldo {
		saque = cc.Saldo - saque
		println("Saldo restante: ", saque)

	} else {
		println("Valor indisponivel para saque!")
	}

}

func (cc *ContaCorrente) transferir() float64 {
	var valorDaTransferencia float64
	var contaDestino *ContaCorrente

	if valorDaTransferencia < cc.Saldo && valorDaTransferencia > 0 {
		println("Digite o valor a ser transferido")
		fmt.Scan(&(valorDaTransferencia))

		println("Qual a conta destino?")
		fmt.Scan(contaDestino)
		contaDestino = cliente.Cliente()

		cc.Saldo -= valorDaTransferencia
		cc.transferir()

		return valorDaTransferencia

	} else {
		println("Não foi possível realizar a transferencia")
		return cc.Saldo
	}
}