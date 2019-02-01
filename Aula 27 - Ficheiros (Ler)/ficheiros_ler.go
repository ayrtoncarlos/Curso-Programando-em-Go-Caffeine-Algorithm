package main

import (
	"fmt"
	"os"
)

var caminho = "Ficheiro.txt"

func lerFicheiro() {

	var ficheiro, _ = os.OpenFile(caminho, os.O_RDONLY, 0644)
	defer ficheiro.Close()

	var conteudo = make([]byte, 1024)

	ficheiro.Read(conteudo)

	fmt.Println(string(conteudo))
}

func main() {

	var _, ficheiro = os.Stat(caminho)

	if(!os.IsNotExist(ficheiro)) {
		lerFicheiro()
	}
}