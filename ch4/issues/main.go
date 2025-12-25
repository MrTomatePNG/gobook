package main

import (
	"fmt"
	"gobook/ch4/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %.11s \n", item.Number, item.Login, item.Title, github.DateFormat(item.CreatedAt))
	}
}
