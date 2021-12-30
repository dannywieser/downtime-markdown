package itunes

import (
	"dgw/downtime/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (
	limit      = "1" // default to one result - use additional search terms to narrow the result
	timeout    = time.Second * 20
	searchPath = "https://itunes.apple.com/search"
)

func DoSearch(title string, media string, author string) SearchResult {
	params := make(map[string]string)
	params["media"] = media
	if author != "" {
		params["attribute"] = "authorTerm"
		params["term"] = author
	} else {
		params["term"] = title
	}
	params["limit"] = limit
	responseBody := utils.DoGet(searchPath, params)

	var response SearchResultWrapper
	jsonErr := json.Unmarshal(responseBody, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var result SearchResult
	if response.ResultCount > 0 {
		fmt.Printf("  ✅ iTunes Search API\n")
		result = response.Results[0]
		parseResult(&result)
	} else {
		fmt.Printf("  ❌ iTunes Search API\n")
	}

	return result
}
