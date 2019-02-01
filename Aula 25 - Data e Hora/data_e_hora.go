package main

import (
	"fmt"
	"time"
)

func main() {

	tempo := time.Now()

	fmt.Println(tempo)

	fmt.Println("Dia:", tempo.Day())
	fmt.Println("Mês:", tempo.Month())
	fmt.Println("Ano:", tempo.Year())

	fmt.Println("Hora:", tempo.Hour())
	fmt.Println("Minuto:", tempo.Minute())
	fmt.Println("Segundo:", tempo.Second())

	data := time.Now().Format("02/01/2006")

	fmt.Println("Data:", data)
}