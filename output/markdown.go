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
	sb.WriteString(fmt.Sprintf("# %s\n", result.TrackName))
	sb.WriteString(fmt.Sprintf("%s/%s\n\n", tagPrefix, result.MediaType))
	if result.Description != "" {
		sb.WriteString(fmt.Sprintf("## Synopsis\n%s\n\n", result.Description))
	}

	sb.WriteString("\n## Author\n")
	if len(result.Artists) > 0 {
		for _, artist := range result.Artists {
			sb.WriteString(fmt.Sprintf("%s\n", buildTag("authors", artist)))
		}
	}

	sb.WriteString("\n## Genre\n")
	for _, genre := range result.Genres {
		if genre != "Books" {
			sb.WriteString(fmt.Sprintf("%s\n", buildTag("genres", genre)))
		}
	}

	sb.WriteString(fmt.Sprintf("- - - -\n ![](%s)\n", result.ArtworkUrl))
	return sb.String()
}

func formatFileName(title string) string {
	filename := title
	filename = strings.Replace(title, " ", "_", -1)
	return fmt.Sprintf("%s.md", strings.ToLower(filename))
}

func SaveToFile(formatted string, title string) {
	err := os.WriteFile(formatFileName(title), []byte(formatted), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
