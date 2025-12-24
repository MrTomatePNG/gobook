package ex43

// Reecreva reverse usando um ponteiro para o array em vez de uma fatia.

func reverse(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}
