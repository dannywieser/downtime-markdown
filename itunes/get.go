package itunes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	limit   = "1" // default to one result - use additional search terms to narrow the result
	timeout = time.Second * 2
)

func BuildQueryParams(req *http.Request, search SearchParams) {
	q := req.URL.Query()
	q.Add("media", search.MediaType)
	q.Add("term", search.Title)
	q.Add("limit", limit)
	req.URL.RawQuery = q.Encode()
}

func DoSearch(search SearchParams) []SearchResult {
	client := http.Client{Timeout: timeout}

	req, err := http.NewRequest(http.MethodGet, apiPath, nil)
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
	if search.Debug {
		fmt.Println(string(body))
	}
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response SearchResultWrapper
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return response.Results
}
