package main

import (
	"dgw/downtime/books"
	"dgw/downtime/config"
	"dgw/downtime/utils"
	"flag"
	"fmt"
)

func main() {
	var title = flag.String("title", "", "Title to search for")
	var author = flag.String("author", "", "Author")
	var mediaType = flag.String("type", "", "books|movies|tv|apps")
	var debug = flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	config.DebugMode = *debug

	markdown := ""
	if *mediaType == "books" {
		fmt.Println("🔎 Searching..")
		fmt.Printf(" Title: %s\n", *title)
		fmt.Printf(" Author: %s\n", *author)
		markdown = books.GenerateBookMarkdown(*title, *author)
	}
	utils.SaveToFile(markdown, *title)
}
