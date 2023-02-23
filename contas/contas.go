package contas

import (
	"fmt"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (cc *ContaCorrente) Deposita() {
	var deposita float64

	println("Digite o valor a ser depositado: ")
	fmt.Scan(&(deposita))
	if deposita > 0 {
		cc.Saldo = deposita + cc.Saldo
		fmt.Printf("Valor após o deposito é de: %g", cc.Saldo)
	} else {
		fmt.Println("Valor de depósito menor que 0!")
	}
}

func (cc *ContaCorrente) Saque() {
	var saque float64

	println("Digite o valor para saque:")
	fmt.Scan(&(saque))

	if saque > 0 && saque <= cc.Saldo {
		saque = cc.Saldo - saque
		fmt.Printf("Saldo restante: %g", saque)

	} else {
		println("Valor indisponivel para saque!")
	}

}

func (cc *ContaCorrente) Transferir() float64 {
	var valorDaTransferencia float64
	var contaDestino int		

	println("Digite a conta para a transferencia")
	fmt.Scan(&(contaDestino))
	println("Digite o valor a ser transferido")
	fmt.Scan(&(valorDaTransferencia))

	if valorDaTransferencia < cc.Saldo {
		fmt.Printf("Transferencia realizada no valor de %g ", valorDaTransferencia)
		println(" para a conta ", contaDestino) 
		return valorDaTransferencia
	} else {
		println("Não foi possível realizar a transferência")
		return cc.Saldo
	}
}
func Cliente() *ContaCorrente {
	cc := &ContaCorrente{}

	fmt.Println("Digite o nome do titular")
	fmt.Scan(&(cc.Titular))

	println("Entre com o número da agencia")
	fmt.Scan(&(cc.NumeroAgencia))

	println("Entre com o número da conta")
	fmt.Scan(&(cc.NumeroConta))

	println("Entre com o saldo")
	fmt.Scan(&(cc.Saldo))

	fmt.Printf("Saldo: %g \n", cc.Saldo)

	return (*ContaCorrente)(cc)
}

func LeComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println("Opção digita foi: ", comando)
	return comando
}

