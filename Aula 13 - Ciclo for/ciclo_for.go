package main

import "fmt"

func main(){

	materialEscolar := [...]string {"Caderno", "Lápis", "Borracha", "Apontador", "Estojo", "Mochila", "Tesoura"}

	/*for indice := 0; indice < len(materialEscolar); indice++ {

		fmt.Printf("materialEscolar[%d]: %s\n", indice, materialEscolar[indice])
	}*/

	for indice, material := range(materialEscolar) {

		fmt.Printf("materialEscolar[%d]: %s\n", indice, material)
	}

	contador := 1

	for contador <= 10 {

		fmt.Printf("Contador: %d\n", contador)
		contador++
	}
}