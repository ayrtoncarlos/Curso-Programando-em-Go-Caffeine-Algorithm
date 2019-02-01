package main

import "fmt"

func main() {

	var x int = 24 // 00011000
	var y int = 37 // 00100101
	// 			      --------
	//	24 & 37 = 0 = 00000000
	fmt.Println("O resultado do operador &:", (x & y))
	// 24 | 37 = 61 = 00111101 
	fmt.Println("O resultado do operador |:", (x | y))
	// 24 << 2 = 96 = 01100000
	fmt.Println("O resultado do operador <<:", (x << 2))
	//  24 >> 2 = 6 = 00000110
	fmt.Println("O resultado do operador >>:", (x >> 2))
	// 24 ^ 37 = 61 = 00111101
	fmt.Println("O resultado do operador ^:", (x ^ y))
   // 24 &^ 37 = 24 = 00011000
	fmt.Println("O resultado do operador &^:", (x &^ y))
  // 24 |^ 37 = -38 = 1111111111111111111111111111111111111111111111111111111111011010
	fmt.Println("O resultado do operador |^:", (x |^ y))
}