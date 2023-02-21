package cliente
import (
	"fmt"
	"home/vitoria/workspace/go/src/banco/contas"

	
)
	

func Cliente() *contas.ContaCorrente {
	cc := &contas.ContaCorrente{}

	fmt.Println("Digite o nome do titular")
	fmt.Scan(&(cc.Titular))

	println("Entre com o número da agencia")
	fmt.Scan(&(cc.NumeroAgencia))

	println("Entre com o número da conta")
	fmt.Scan(&(cc.NumeroConta))

	println("Entre com o saldo")
	fmt.Scan(&(cc.Saldo))

	fmt.Printf("Saldo: %g \n", cc.Saldo)

	return (*contas.ContaCorrente)(cc)
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println("Opção digita foi: ", comando)
	return comando
}