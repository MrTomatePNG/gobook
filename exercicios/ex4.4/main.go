package ex44

// Exreva um versao de rotate que funcione com um unico passo.

func rotate(arr *[]int) {
	n := len(*arr)
	if n == 0 {
		return
	}
	first := (*arr)[0]
	for i := 0; i < n-1; i++ {
		(*arr)[i] = (*arr)[i+1]
	}
	(*arr)[n-1] = first
}
