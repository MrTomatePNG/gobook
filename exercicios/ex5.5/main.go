package ex55

import (
	"strings"

	"golang.org/x/net/html"
)

// Implemente contWordsAndImage(n *html.Node) (word, images int)
// (Veja o exercicio 4.9 para saber como separar palavras)

func countWordsAndImgages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childWords, childImages := countWordsAndImgages(c)
		words += childWords
		images += childImages
	}
	return words, images
}
