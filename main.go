package main

import (
	"dgw/downtime/config"
	itunes "dgw/downtime/itunes"
	output "dgw/downtime/output"
	"flag"
	"fmt"
)

func main() {
	var title = flag.String("title", "", "Title to search for")
	var mediaType = flag.String("type", "", "books|movies|tv|apps")
	var debug = flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	params := itunes.SearchParams{
		Title:     *title,
		MediaType: *mediaType,
	}
	config.DebugMode = *debug

	result := itunes.DoSearch(params)
	markdown := output.FormatResult(result)
	fmt.Print(markdown)
	output.SaveToFile(markdown, *title)
}
