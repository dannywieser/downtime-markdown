package output

import (
	itunes "dgw/downtime/itunes"
	"fmt"
	"log"
	"os"
	"strings"
)

func h2(title string) string {
	return fmt.Sprintf("%s %s\n", h2Prefix, title)
}

func buildTag(tagType string, tagValue string) string {
	return fmt.Sprintf("%s/%s/%s#", tagPrefix, tagType, tagValue)
}

func FormatArtist(result itunes.SearchResult) string {
	var sb strings.Builder

	artistLabel := artistTitleMap[result.MediaType]
	artistTag := fmt.Sprintf("%ss", strings.ToLower(artistLabel))
	sb.WriteString(h2(artistLabel))
	if len(result.Artists) > 0 {
		for _, artist := range result.Artists {
			sb.WriteString(fmt.Sprintf("%s\n", buildTag(artistTag, artist)))
		}
	} else {
		sb.WriteString(fmt.Sprintf("%s\n", buildTag(artistTag, result.ArtistName)))
	}
	return sb.String()
}

func FormatGenre(result itunes.SearchResult) string {
	var sb strings.Builder
	sb.WriteString(h2("Genre"))

	if len(result.Genres) > 0 {
		for _, genre := range result.Genres {
			if genre != "Books" {
				sb.WriteString(fmt.Sprintf("%s\n", buildTag(genreTag, genre)))
			}
		}
		return sb.String()
	}

	if result.PrimaryGenre != "" {
		sb.WriteString(fmt.Sprintf("%s\n", buildTag(genreTag, result.PrimaryGenre)))
		return sb.String()
	}

	sb.WriteString("Not Available\n")
	return sb.String()
}

func FormatImage(result itunes.SearchResult) string {
	return fmt.Sprintf("\n- - - -\n ![](%s)\n", result.ArtworkUrl)
}

func FormatResult(result itunes.SearchResult) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s\n", result.TrackName))
	sb.WriteString(fmt.Sprintf("%s/%s\n\n", tagPrefix, result.MediaType))

	sb.WriteString(FormatSynopsis(result))
	sb.WriteString(FormatArtist(result))
	sb.WriteString(FormatGenre(result))
	sb.WriteString(FormatImage(result))

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
