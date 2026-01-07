package links

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "MyGoCrawler/1.0")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {

		// Dentro do Extract, no visitNode
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, _ := resp.Request.URL.Parse(a.Val)

					// Filtro: só segue links que parecem ser de páginas de wallpapers
					if strings.Contains(link.Path, "/wallpapers/") && !strings.HasSuffix(link.Path, ".png") {
						links = append(links, link.String())
					}
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}
func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}

}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
func downloadImage(url string) error {
	// 1. Faz a requisição
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 2. Cria o arquivo local usando o nome final da URL
	fileName := filepath.Base(url)
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// 3. Copia os bytes da resposta diretamente para o arquivo
	_, err = io.Copy(out, resp.Body)
	return err
}
func saveImage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Cria o nome do arquivo a partir da URL
	fileName := filepath.Base(url)
	fmt.Printf("Baixando: %s\n", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
