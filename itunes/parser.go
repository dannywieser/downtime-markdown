package itunes

import (
	"strings"
)

var htmlToMarkDownMap = map[string]string{
	"<b>":    "**",
	"</b>":   "**",
	"<i>":    "*",
	"</i>":   "*",
	"<br />": "\n",
}

func replaceHtmlWithMarkdown(stringToParse string) string {
	updatedString := stringToParse
	for htmlString, markdownReplacement := range htmlToMarkDownMap {
		updatedString = strings.Replace(updatedString, htmlString, markdownReplacement, -1)
	}
	return updatedString
}

func parseResult(result *SearchResult, search SearchParams) {
	result.Description = replaceHtmlWithMarkdown(result.Description)
	result.ArtworkUrl = strings.Replace(result.ArtworkUrl, "100x100", "200x200", 1)
	for _, artistId := range result.ArtistIds {
		artist := LookupArtist(artistId, false)
		result.Artists = append(result.Artists, artist.Name)
	}
	result.MediaType = search.MediaType
}
