package main

import "fmt"

// NOTA: exemplos e uma breve explciação sobre
// cada operador BIT a BIT
func main() {
	// valores iniciais
	// observe os valores em binario
	a := 5 // 0101
	b := 3 // 0011

	// AND
	// o operador `&` retorna 1 se e somente se
	// ambos os valores de entrada forem 1
	andResult := a & b //0001 = 1
	fmt.Printf("%04b\t=%d\n", andResult, andResult)
	// OR
	// o operador `|`  retorna 1 se *pelo menos*
	// um de seus valores de entrada for 1
	orResult := a | b // 0111 = 7
	fmt.Printf("%04b \t=%d\n", orResult, orResult)
}
