package main

import (
	"fmt"
	"gobook/ch5/links"
	"log"
	"os"
	"sync"
)

func main() {
	worklist := make(chan []string)
	var wg sync.WaitGroup

	wg.Go(func() {
		worklist <- os.Args[1:]
	})

	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				wg.Add(1)
				go func(link string) {
					defer wg.Done()
					worklist <- crawl(link)
				}(link)
			}
		}
	}

	go func() {
		wg.Wait()
		close(worklist)
	}()
}

func crawl(link string) []string {
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}
	return list
}
