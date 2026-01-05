package ex57

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Desnvolva startElement e endElement em um pretty printer HTML
// generico. Exiba nós de texto e os atributos de cada elemento
// (<a href='...'). Use formas compactas como </img> no lugar de <img> </img>
// quando um elemento não tiver filhos. Escreva um teste para garantir que
// seja póssivel fazer parse da saída com sucesso (VEJA o capitulo 11)

func PrettyPrint(n *html.Node, w io.Writer) {
	forEachNode(n, w, 0, startElement, endElement)
}

func forEachNode(n *html.Node, w io.Writer, depth int, pre, post func(n *html.Node, w io.Writer, depth int)) {
	if pre != nil {
		pre(n, w, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, w, depth+1, pre, post)
	}
	if post != nil {
		post(n, w, depth)
	}
}

func startElement(n *html.Node, w io.Writer, depth int) {
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(w, "%*s<%s", depth*2, "\t", n.Data)
		for _, attr := range n.Attr {
			fmt.Fprintf(w, " %s='%s'", attr.Key, attr.Val)
		}

		if n.FirstChild == nil {
			fmt.Fprintf(w, "/>\n")
		} else {
			fmt.Fprintf(w, ">\n")
		}
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Fprintf(w, "%*s%s\n", depth*2, "", text)
		}
	}
}

func endElement(n *html.Node, w io.Writer, depth int) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
