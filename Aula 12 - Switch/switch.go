package main

import "fmt"

func main(){

	caso := "E"

	switch(caso){

	case "A":
		fmt.Println("O caso A existe.")
	case "B": 
		fmt.Println("O caso B existe.")
	case "C":
		fmt.Println("O caso C existe.")
	default:
		fmt.Println("O caso", caso, "não existe.")
	}
}