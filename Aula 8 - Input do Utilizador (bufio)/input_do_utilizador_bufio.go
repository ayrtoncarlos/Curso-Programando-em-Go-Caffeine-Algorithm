package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

	inputDoUtilizador := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o seu primeiro nome: ")
	primeiroNome, _ := inputDoUtilizador.ReadString('\n')

	fmt.Print("Digite o seu último nome: ")
	ultimoNome, _ := inputDoUtilizador.ReadString('\n')

	fmt.Print("Digite a sua idade: ")
	idade, _ := inputDoUtilizador.ReadString('\n')

	fmt.Printf("Primeiro nome: %sÚltimo nome: %sIdade: %s", primeiroNome, ultimoNome, idade)
}