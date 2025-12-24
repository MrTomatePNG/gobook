package main

import "fmt"

func ehVendedorDeManga(name string) bool {
	return len(name) > 0 && name[len(name)-1] == 'm'
}

func main() {
	grafo := make(map[string][]string)
	grafo["voce"] = []string{"alice", "bob", "claire"}
	grafo["bob"] = []string{"anuj", "peggy"}
	grafo["alice"] = []string{"peggy"}
	grafo["claire"] = []string{"thom", "jonny"}
	grafo["anuj"] = []string{}
	grafo["peggy"] = []string{}
	grafo["thom"] = []string{}
	grafo["jonny"] = []string{}

	fila := grafo["voce"]

	verificados := make(map[string]bool)

	for len(fila) > 0 {
		pessoa := fila[0]
		fila = fila[1:]
		if !verificados[pessoa] {
			if ehVendedorDeManga(pessoa) {
				fmt.Printf("%s é vendor de manga !\n", pessoa)
				return
			} else {
				fila = append(fila, grafo[pessoa]...)
				verificados[pessoa] = true
			}
		}
	}

	fmt.Println("Nenhum vendedor de manga encontrado")
}
