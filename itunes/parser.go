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

func getSynopsis(result SearchResult) string {
	if result.Description != "" {
		return replaceHtmlWithMarkdown(result.Description)
	}

	if result.LongDescription != "" {
		return replaceHtmlWithMarkdown(result.LongDescription)
	}

	return "Not Available"
}

func parseResult(result *SearchResult) {
	// TODO: artists?
	result.ArtworkUrl = strings.Replace(result.ArtworkUrl, "100x100", "200x200", 1)
	result.Synopsis = getSynopsis(*result)
}
