package main

import "fmt"

/*
	Existe três tipos de erros: erros e compilação, erros de tempo de execução e erros de lógica;
	Os erros de compilação são erros que impedem a execução da nossa aplicação e são maioritariamente causados pelo mal digitar do nosso código
	(por exemplo a não inserção do ';' para finalizar uma determinada linha de código);
	Os erros de tempo de execução são erros que ocorrem enquanto a nossa aplicação é executada e normalmente damos conta desses mesmos quando a
	aplicação tenta uma operação que é impossível executar (por exemplo a divisão por zero);
	Os erros de lógica são erros que impedem que a nossa aplicação faça o que é pretendido fazer, ou seja, o nosso código é compilado e
	executado sem erros, mas o resultado de uma operação acaba por não ser o esperado (por exemplo não complementar uma dada variável do tipo string);
	Destes três tipos de erros, os erros lógicos são os mais dificeis de localizar e corrigir.
*/
func main(){

	fmt.Println("Os três tipos de erros comuns:\n1º) Erros de compilação;\n2º) Erros de tempo de execução;\n3º) Erros de lógica.")
}