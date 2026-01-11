package main

import (
	"fmt"
	"strings"
)

// isso aqui é legal
type FuncStruct struct {
	Number int
	Op     func(number int) int
}

func main() {
	f := square
	fmt.Println(f(2))
	f = negative
	fmt.Println(f(2))
	fmt.Printf("%T\n", f)
	fmt.Println(mapper(func(n int) int {
		return n
	}))

	fmt.Println(strings.Map(add1, "HAL-9000"))
	adica := func(n int) int {
		return 1 + n
	}

	fs := FuncStruct{
		Number: 1,
		Op:     adica,
	}
	result := fs.Op(1)

	fmt.Println(result)

}
func square(n int) int {
	return n * n
}
func negative(n int) int {
	return -n
}

func product(x, y int) int {
	return x * y
}

func mapper[T comparable](f func(int) T) bool {
	if f(2) == f(3) {
		return true
	}
	return false
}
func add1(r rune) rune { return r + 1 }
