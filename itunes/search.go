package itunes

import (
	"dgw/downtime/config"
	"dgw/downtime/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	limit      = "1" // default to one result - use additional search terms to narrow the result
	timeout    = time.Second * 20
	searchPath = "https://itunes.apple.com/search"
)

func BuildQueryParams(req *http.Request, title string, media string) {
	queryLimit := limit
	if config.DebugMode {
		queryLimit = "1"
	}
	q := req.URL.Query()
	q.Add("media", media)
	q.Add("term", title)
	q.Add("attribute", "titleTerm")

	q.Add("limit", queryLimit)
	req.URL.RawQuery = q.Encode()
}

func DoSearch(title string, media string) SearchResult {
	params := make(map[string]string)
	params["media"] = media
	params["term"] = title
	//params["attribute"] = "titleTerm"
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
