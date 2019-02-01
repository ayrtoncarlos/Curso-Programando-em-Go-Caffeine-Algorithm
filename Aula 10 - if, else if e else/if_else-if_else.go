package main

import "fmt"

/*
	(if - se) Condição for verdadeira {
		O código dentro do if é executado.
	} (else if - senão se) Condição for verdadeira (Só ocorre o else if caso a condição de if seja falsa) {
		O código dentro do else if é executado.
	} (else - senão) Sem condição (Só ocorre caso a condição de if e de else if sejam falsas) {
		O código dentro do else é executado.
	}
*/
	
func main(){

	x := 30

	if(x == 10){
		fmt.Println("O valor de x é igual a 10.")
	} else if(x == 20){
		fmt.Println("O valor de x é igual a 20.")
	} else{
		fmt.Println("O valor de x é diferente de 10 e de 20.")
	}
}