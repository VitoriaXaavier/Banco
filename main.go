package main

import (
	"fmt"
	"os"
	"/home/vitoria/workspace/go/src/banco/cliente"

)


func menu() {
	println("Deseja realizar alguma operação?")
	println("1- Depositar")
	println("2- Sacar")
	println("3- Transferir")
	println("4- sair")

}

func main() {

	cc := cliente.Cliente()

	menu()

	comando := cliente.leComando()

	switch comando {
	case 1:
		cc.deposita()
	case 2:
		cc.saque()
	case 3:
		cc.transferir()
	case 4:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		println("Opção invalida!")
		os.Exit(-1)

	}

}
