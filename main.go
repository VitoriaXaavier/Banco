package main

import (
	"fmt"
	"os"
	"banco/cliente"

)


func Menu() {
	println("Deseja realizar alguma operação?")
	println("1- Depositar")
	println("2- Sacar")
	println("3- Transferir")
	println("4- sair")

}

func Main() {

	cc := cliente.Cliente()

	Menu()

	comando := cliente.leComando()

	switch comando {
	case 1:
		cc.Deposita()
	case 2:
		cc.Saque()
	case 3:
		cc.Transferir()
	case 4:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		println("Opção invalida!")
		os.Exit(-1)

	}

}
