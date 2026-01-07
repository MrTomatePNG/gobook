package main

import "fmt"

func matchAPlusB(s string) bool {
	if len(s) == 0 {
		return false
	}

	state := 0
	for i := 0; i < len(s); i++ {
		char := s[i]

		switch state {
		case 0:
			if char == 'a' {
				state = 1
			} else {
				return false
			}
		case 1:
			if char == 'a' {
				state = 1
			} else if char == 'b' {
				state = 2
			} else {
				return false
			}
		case 2:
			return false
		}
	}

	return state == 2
}

func main() {
	tests := []string{"ab", "aaaab", "b", "a", "aaac", "aaab "}

	for _, t := range tests {
		fmt.Printf("'%s': %v\n", t, matchAPlusB(t))
	}
}
