package main

import "fmt"

type Estrutura struct {
	numero int
}

func (estrutura Estrutura) imprimirNumero() {

	if(estrutura.numero >= 0 && estrutura.numero <= 100) {

		fmt.Println("Número:", estrutura.numero)
	} else {

		fmt.Println("O número precisa de ser >= 0 e <= 100.")
	}
}

func main(){

	estruturaA := Estrutura { 100 }
	estruturaB := Estrutura {}

	estruturaA.imprimirNumero()

	estruturaB.numero = 200

	estruturaB.imprimirNumero()
}