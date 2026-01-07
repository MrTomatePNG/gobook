package main

import (
	"fmt"
	"gobook/ch5/links"
	"log"
	"os"
	"strings"
)

func breadFirsth(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	breadFirsth(crawl, os.Args[1:])
}
