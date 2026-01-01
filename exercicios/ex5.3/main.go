package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Escreva uma função para exibir o conteúdo de todos
// os nós de texto em uma árvore de documento HTML.
// Não desça em elementos <script> ou <style>, pois seus conteúdos
// não são visivéis em um navegador navegador web

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}
	printTextNodes(doc)
}

func printTextNodes(n *html.Node) {

	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Println(text)
		}
	}

	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" || n.Data == "header" {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printTextNodes(c)
	}
}
