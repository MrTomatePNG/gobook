package main

import "fmt"

func romanToInt(s string) int {
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0

	for i := 0; i < len(s); i++ {
		currentValue := romanToIntMap[rune(s[i])]

		if i+1 < len(s) {
			nextValue := romanToIntMap[rune(s[i+1])]
			if currentValue < nextValue {
				total -= currentValue
			} else {
				total += currentValue
			}
		} else {
			total += currentValue
		}
	}
	return total
}

func main() {
	fmt.Println(romanToInt("III"))     // Saída: 3
	fmt.Println(romanToInt("IV"))      // Saída: 4
	fmt.Println(romanToInt("IX"))      // Saída: 9
	fmt.Println(romanToInt("LVIII"))   // Saída: 58
	fmt.Println(romanToInt("MCMXCIV")) // Saída: 1994
}
