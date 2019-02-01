package main

import "fmt"

func main(){

	animais := [...]string {"Cachorro", "Gato", "Galinha", "Coelho", "Leão"}

	for indice, animal := range(animais) {

		fmt.Printf("animais[%d]: %s\n", indice, animal)

		if(animal == "Galinha") {
			break
		}
	}

	contador := 0

	for(contador < 10) {

		contador++

		if(contador == 5) {
			continue
		}

		fmt.Printf("Contador: %d\n", contador)
	}
}