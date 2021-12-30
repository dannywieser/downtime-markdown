package main

import (
	"dgw/downtime/books"
	"dgw/downtime/config"
	"flag"
	"fmt"
)

func main() {
	var title = flag.String("title", "", "Title to search for")
	var mediaType = flag.String("type", "", "books|movies|tv|apps")
	var debug = flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	config.DebugMode = *debug

	markdown := ""
	if *mediaType == "books" {
		markdown = books.GenerateBookMarkdown(*title)
	}

	fmt.Printf("\n\n======== Results ==========\n\n")
	fmt.Print(markdown)
}
