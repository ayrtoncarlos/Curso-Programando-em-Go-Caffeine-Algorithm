package main

import "fmt"

var acesso = "Global"

func main(){
	mudarAcesso()
	fmt.Println("Acesso no exterior da função:", acesso)
}

func mudarAcesso(){

	var acesso = "Local"
	fmt.Println("Acesso no interior da função:", acesso)
}