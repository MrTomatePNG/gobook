package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reverso := 0
	for x != 0 {
		digit := x % 10
		reverso = reverso*10 + digit
	}
	x /= 10
	return original == reverso

}
