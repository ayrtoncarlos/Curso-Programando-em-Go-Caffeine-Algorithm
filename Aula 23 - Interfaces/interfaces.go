package main

import (
	"fmt"
	"math"
)

type FiguraGeometrica interface {

	area() float64
	perimetro() float64
}

type Retangulo struct {

	largura, altura float64
}

type Circulo struct {

	raio float64
}

func (retangulo Retangulo) area() float64 {

	return (retangulo.altura * retangulo.largura)
} 

func (retangulo Retangulo) perimetro() float64 {

	return ((2 * retangulo.altura) + (2 * retangulo.largura))
}

func (circulo Circulo) area() float64 {

	return (math.Pi * (circulo.raio * circulo.raio))
}

func (circulo Circulo) perimetro() float64 {

	return (2 * (math.Pi * circulo.raio))
}

func calcularArea(figuraGeometrica FiguraGeometrica) {

	fmt.Println("\tArea:", figuraGeometrica.area())
	fmt.Println("\tPerímetro:", figuraGeometrica.perimetro())
}
func main() {

	retagulo := Retangulo { 3, 4 }
	circulo := Circulo { 5 }

	fmt.Println("Retângulo:")
	calcularArea(retagulo)

	fmt.Println("Círculo:")
	calcularArea(circulo)
}