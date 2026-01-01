package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Escreva uma função pra preecher um mapeamento de nomes
// de elementos -- p, div, spam, e assim por diante -- para
// o número de elementos com esse nome em uma árvore de documento HTML.

var m = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro: %v\n", err)
		os.Exit(1)
	}

	elementCount := make(map[string]int)
	elementCount = visitMapper(elementCount, doc)

	// Imprimir contagem de elementos
	for name, count := range elementCount {
		fmt.Printf("%s: %d\n", name, count)
	}
}

func visitMapper(m map[string]int, n *html.Node) map[string]int {

	if n == nil {
		return m
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	m = visitMapper(m, n.FirstChild)
	m = visitMapper(m, n.NextSibling)
	return m
}
