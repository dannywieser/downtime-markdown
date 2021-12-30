package openlibrary

import (
	"dgw/downtime/utils"
	"encoding/json"
	"fmt"
	"log"
)

const (
	searchPath = "http://openlibrary.org/search.json"
	coverPath  = "https://covers.openlibrary.org/b/id"
)

func DoSearch(title string, author string) OpenLibrarySearchResult {
	params := make(map[string]string)
	params["title"] = title
	if author != "" {
		params["author"] = author
	}
	responseBody := utils.DoGet(searchPath, params)

	var response searchResultWrapper
	jsonErr := json.Unmarshal(responseBody, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var result OpenLibrarySearchResult
	if response.NumFound > 0 {
		result = response.Docs[0] // assume the first result is the closest match
		result.CoverImageUrl = fmt.Sprintf("%s/%d-M.jpg", coverPath, result.CoverId)
		fmt.Printf("  ✅ Open Library API\n")
	} else {
		fmt.Printf("  ❌ Open Library API\n")
	}

	return result
}
