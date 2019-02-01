package main

import "os"

var caminho = "Ficheiro.txt"
var paises = [...]string { "Brasil", "Argentina", "Colombia", "Chile", "Uruguai", "Peru", "Bolívia", "Venezuela"}

func escreverFicheiro() {

	var ficheiro, _ = os.OpenFile(caminho, os.O_WRONLY, 0644)
	defer ficheiro.Close()

	for _, pais := range(paises) {

		ficheiro.WriteString(pais + "\n")
	}

	ficheiro.Sync()
}

func main() {

	var _, ficheiro = os.Stat(caminho)

	if(!os.IsNotExist(ficheiro)) {
		escreverFicheiro()
	}
}