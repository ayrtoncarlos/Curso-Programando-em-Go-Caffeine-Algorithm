package main

import "fmt"

func main(){

	//var cores [5]string

	/*cores[0] = "Preto"
	cores[1] = "Amarelo"
	cores[2] = "Verde"
	cores[3] = "Azul"
	cores[4] = "Vermelho"*/

	cores := [...]string {"Preto", "Amarelo", "Verde", "Azul", "Vermelho"}

	fmt.Println("Número de cores:", len(cores))
	fmt.Println("Primeira cor:", cores[0])
	fmt.Println("Última cor:", cores[len(cores) - 1])
}