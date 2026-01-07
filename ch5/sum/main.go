package main

import "fmt"

func sum(vals ...int) int {
	var total int
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	values := []int{1, 2, 3}
	fmt.Println(sum(values...))
}
