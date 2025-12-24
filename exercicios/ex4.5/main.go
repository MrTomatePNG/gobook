package ex45

// Escreva uma função in-place para eliminar duplicatas em adjacentes em uma fatia de strings.

func removeAdjacentDuplicates(strings *[]string) {
	n := len(*strings)
	if n == 0 {
		return
	}

	j := 0
	for i := 1; i < n; i++ {
		if (*strings)[i] != (*strings)[j] {
			j++
			(*strings)[j] = (*strings)[i]
		}
	}
	*strings = (*strings)[:j+1]
}
