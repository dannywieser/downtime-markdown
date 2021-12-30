package openlibrary

import (
	"dgw/downtime/config"
	"dgw/downtime/model"
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
	searchPath = "http://openlibrary.org/search.json"
)

func DoSearch(search model.SearchParams) book {
	client := http.Client{Timeout: timeout}

	req, err := http.NewRequest(http.MethodGet, searchPath, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("q", fmt.Sprintf("title:%s", search.Title))
	q.Add("fields", "title,author_name,isbn")
	req.URL.RawQuery = q.Encode()

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

	var response searchResultWrapper
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if response.NumFound == 0 {
		log.Fatal(fmt.Errorf("no results found for title [%s] and type [%s]", search.Title, search.MediaType))
		os.Exit(1)
	}
	return getBook(response.Docs[0].ISBN[0])
}
