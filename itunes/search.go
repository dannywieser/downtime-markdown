package itunes

import (
	"dgw/downtime/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	limit      = "1" // default to one result - use additional search terms to narrow the result
	timeout    = time.Second * 20
	searchPath = "https://itunes.apple.com/search"
)

var typeMap = map[string]string{
	"books":  "ebook",
	"movies": "movie",
}

func BuildQueryParams(req *http.Request, search SearchParams) {
	queryLimit := limit
	if config.DebugMode {
		queryLimit = "1"
	}
	q := req.URL.Query()
	q.Add("media", search.MediaType)
	q.Add("term", search.Title)
	q.Add("attribute", "titleTerm")

	q.Add("limit", queryLimit)
	req.URL.RawQuery = q.Encode()
}

func DoSearch(search SearchParams) SearchResult {
	client := http.Client{Timeout: timeout}

	req, err := http.NewRequest(http.MethodGet, searchPath, nil)
	if err != nil {
		log.Fatal(err)
	}
	BuildQueryParams(req, search)

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if config.DebugMode {
		fmt.Printf("\n=== Search ====\n%s\n%s", req.URL, string(body))

	}
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response SearchResultWrapper
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if response.ResultCount == 0 {
		log.Fatal(fmt.Errorf("no results found for title [%s] and type [%s]", search.Title, search.MediaType))
		os.Exit(1)
	}

	if config.DebugMode {
		fmt.Printf("results: %d\n", response.ResultCount)
		// fmt.Printf("\n=== Search ====\n%s\n%s", req.URL, string(body))
		for _, debugResult := range response.Results {
			fmt.Printf("Title: %s\n", debugResult.TrackName)
		}
	}

	result := response.Results[0]
	parseResult(&result, search)
	return result
}
