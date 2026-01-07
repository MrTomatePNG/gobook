package main

import "fmt"

func Protecao() {
	if n := recover(); n != nil {
		fmt.Printf("recuperamos do erro: %v\n", n)
	}
}

func DivideByZero(a, b int) {
	defer Protecao()
	if b == 0 {
		panic("divisao por zero")
	}

	fmt.Println("Resultado:", a/b)
}

func main() {
	DivideByZero(10, 2)
	DivideByZero(10, 0) // Isso causaria um crash, mas o recover vai salvar
	fmt.Println("O programa continuou normalmente!")
	DivideByZero(30, 20)
}
