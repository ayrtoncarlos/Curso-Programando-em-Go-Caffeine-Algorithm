package main

import "fmt"

var pessoas = map[string]int {"Ayrton Andrade": 24,
							  "Vanessa Araújo": 26,
							  "Cristine Diniz": 30}
func main(){

	pessoas["Felipe Malheiros"] = 26
	pessoas["Cassio Klebio"] = 40

	pessoas["Felipe Malheiros"] = 27
	delete(pessoas, "Cassio Klebio")
	
	fmt.Println("Número de pessoas:", len(pessoas))

	for chave, valor := range(pessoas) {

		fmt.Println("Nome:", chave)
		fmt.Println("Idade:", valor)
	}
	pessoas = make(map[string]int) // Zerar o dicionario.
}