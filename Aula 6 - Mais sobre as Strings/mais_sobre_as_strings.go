package main

import (
	"fmt"
	"strings"
)

func main(){

	primeiroNome, ultimoNome := "Ayrton", "Andrade"
	idade := 24

	fmt.Println("Nome:", primeiroNome, ultimoNome, "\nIdade:", idade)
	fmt.Println("Nome:", len(primeiroNome), len(ultimoNome), "\nIdade:", idade)
	fmt.Printf("Nome: %s %s\nIdade: %d", strings.ToUpper(primeiroNome), strings.ToLower(ultimoNome), idade)
}