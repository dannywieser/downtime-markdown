package output

import (
	itunes "dgw/downtime/itunes"
	"fmt"
	"strings"
)

func FormatResult(result itunes.SearchResult) {
	fmt.Printf("# %s\n\n", result.TrackName)
	fmt.Printf("## Synopsis\n%s\n\n", result.Description)
	fmt.Printf("## Author\n%s\n\n", result.ArtistName)
	fmt.Printf("## Genres11\n%s\n\n", strings.Join(result.Genres, "\n"))
	fmt.Printf("image: %s", result.ArtworkUrl)
}
