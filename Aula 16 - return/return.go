package main

import "fmt"

func main(){

	fmt.Println("O resultado da adição:", adicao())
}

func adicao() int {

	resultado := 0

	for valor := 0; valor <= 10; valor++ {

		resultado += valor
	}

	return resultado
}