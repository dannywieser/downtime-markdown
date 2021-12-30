package output

import (
	itunes "dgw/downtime/itunes"
	"fmt"
	"log"
	"os"
	"strings"
)

const tagPrefix = "#downtime"

func buildTag(tagType string, tagValue string) string {
	return fmt.Sprintf("%s/%s/%s#", tagPrefix, tagType, tagValue)
}

func FormatResult(result itunes.SearchResult) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s\n\n", result.TrackName))
	if result.Description != "" {
		sb.WriteString(fmt.Sprintf("## Synopsis\n%s\n\n", result.Description))
	}

	sb.WriteString("\n## Author\n")
	if result.ArtistName != "" {
		sb.WriteString(fmt.Sprintf("%s\n", buildTag("authors", result.ArtistName)))
	}

	sb.WriteString("\n## Genre\n")
	for _, genre := range result.Genres {
		sb.WriteString(fmt.Sprintf("%s\n", buildTag("genres", genre)))
	}

	sb.WriteString(fmt.Sprintf("- - - -\n ![](%s)\n", result.ArtworkUrl))
	return sb.String()
}

func SaveToFile(formatted string, filename string) {
	err := os.WriteFile(filename, []byte(formatted), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
