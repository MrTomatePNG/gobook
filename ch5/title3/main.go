package title3

import (
	"fmt"
	"gobook/ch5/outline2"

	"golang.org/x/net/html"
)

// soletitle devolve o texto  do primeiro elemento title nao vazio
// em doc , e um error se não houver exatamanete um.

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			//sem panico
		case bailout{}:
			//panico esperado
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) //panico inesperado ; prossegue com o panico
		}
	}()
	//sai da recursao se mais de um titulo nao vazio for encontrado
	outline2.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title == "" {
				panic(bailout{}) //varios elementos tittle
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no tittle element")
	}
	return title, err
}
