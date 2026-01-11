package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	var i []int = twoSum(numbers, target)
	fmt.Println(i)
}

func twoSum(numbers []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range numbers {
		need := target - v
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return nil
}
