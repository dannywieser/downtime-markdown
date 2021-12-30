package main

import (
	itunes "dgw/downtime/itunes"
	output "dgw/downtime/output"
	"flag"
	"fmt"
)

func main() {
	var title = flag.String("title", "", "Title to search for")
	var mediaType = flag.String("type", "", "book|movie|tv|app")
	var debug = flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	params := itunes.SearchParams{
		Title:     *title,
		MediaType: *mediaType,
		Debug:     *debug,
	}

	result := itunes.DoSearch(params)
	markdown := output.FormatResult(result)
	fmt.Print(markdown)
	output.SaveToFile(markdown, "test.md")
}
