package main

import "fmt"

var arrayMultidimensional = [5][4]int {{1, 2, 3, 4},
									   {1, 1, 1, 1},
									   {2, 2, 2, 2},
									   {3, 3, 3, 3},
									   {4, 4, 4, 4}}
func main(){

	for linha := 0; linha < 5; linha++ {
		
		for coluna := 0; coluna < 4; coluna++ {

			fmt.Print(arrayMultidimensional[linha][coluna], " ")
		}
		fmt.Println()
	}
}