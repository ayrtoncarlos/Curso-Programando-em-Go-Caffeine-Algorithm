package main

import "fmt"

func main(){

	x := 10
	y := 20

	// Operadores de Comparação:
	// == --> Comparação
	// != --> Diferente
	// > --> Maior que
	// < --> Menor que
	// >= --> Maior que ou igual
	// <= --> Menor que ou igual

	// Alguns operadores lógicos:
	// E (&&)
	// V - V = V
	// V - F = F
	// F - V = F
	// F - F = F

	// OU (||)
	// V - V = V
	// V - F = V
	// F - V = V
	// F - F = F

	if(x <= y || x == y){
		fmt.Println("Esta condição é verdadeira.")
	} else{
		fmt.Println("Esta condição é falsa.")
	}
}