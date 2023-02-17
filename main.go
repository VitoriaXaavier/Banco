package main

import (
	"fmt"
	"os"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}


func (cc *ContaCorrente) deposita() {
	var deposita float64

	println("Digite o valor a ser depositado: ")
	fmt.Scan(&(deposita))
	
	cc.saldo = deposita + cc.saldo
	fmt.Println("Valor após o deposito é de: ", cc.saldo)
}

func (cc *ContaCorrente) saque() {
	var saque float64

	println("Digite o valor para saque:")
	fmt.Scan(&(saque))

	if saque <= cc.saldo {
		saque = cc.saldo - saque
		println("Saldo restante: ", saque)

	} else {
		println("Valor indisponivel para saque")
	}

}
func cliente() *ContaCorrente {
	cc := &ContaCorrente{}

	fmt.Println("Digite o nome do titular")
	fmt.Scan(&(cc.titular))

	println("Entre com o número da agencia")
	fmt.Scan(&(cc.numeroAgencia))

	println("Entre com o número da conta")
	fmt.Scan(&(cc.numeroConta))

	println("Entre com o saldo")
	fmt.Scan(&(cc.saldo))

	fmt.Printf("Saldo: %g \n", cc.saldo)
	return (*ContaCorrente)(cc)
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println("Opção digita foi: ", comando)
	return comando
}

func menu() {
	println("Deseja realizar alguma operação?")
	println("1- Depositar")
	println("2- Sacar")
	println("3- sair")

}

func main() {

	cc := cliente()

	menu()

	comando := leComando()

	switch comando {
	case 1:
		cc.deposita()
	case 2:
		cc.saque()
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		println("Opção invalida!")
		os.Exit(-1)

	}

}

