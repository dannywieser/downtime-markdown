package main

import (
	itunes "dgw/downtime/itunes"
	"flag"
	"fmt"
	"strings"
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

	results := itunes.DoSearch(params)
	firstResult := results[0]
	fmt.Printf("# %s\n\n", firstResult.TrackName)
	fmt.Printf("## Synopsis\n%s\n\n", firstResult.Description)
	fmt.Printf("## Author\n%s\n\n", firstResult.ArtistName)
	fmt.Printf("## Genres\n%s\n\n", strings.Join(firstResult.Genres, "\n"))
}
