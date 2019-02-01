package main

import "fmt"

func main(){

	dadosPessoais("Ayrton Andrade", 24, "Brasileiro")
}

func dadosPessoais(nome string, idade int, nacionalidade string){

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Nacionalidade:", nacionalidade)
}