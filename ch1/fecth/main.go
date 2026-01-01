package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix("https://", url) {
			url = "http://" + url
		}

		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		dst := os.Stdout

		b, err := io.Copy(dst, res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", b)
	}
}
