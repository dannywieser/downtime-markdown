package main

import (
	"dgw/downtime/config"
	"dgw/downtime/model"
	"dgw/downtime/openlibrary"
	"flag"
)

func main() {
	var title = flag.String("title", "", "Title to search for")
	var mediaType = flag.String("type", "", "books|movies|tv|apps")
	var debug = flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	params := model.SearchParams{
		Title:     *title,
		MediaType: *mediaType,
	}
	config.DebugMode = *debug

	openlibrary.DoSearch(params)

	//result := itunes.DoSearch(params)
	//markdown := output.FormatResult(result)
	//fmt.Print(markdown)
	//output.SaveToFile(markdown, *title)
}
