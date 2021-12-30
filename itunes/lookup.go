package itunes

import (
	"dgw/downtime/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	lookupPath = "https://itunes.apple.com/lookup"
)

func LookupArtist(id int) ArtistResult {
	client := http.Client{Timeout: timeout}

	req, err := http.NewRequest(http.MethodGet, lookupPath, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", fmt.Sprintf("%d", id))
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
		fmt.Printf("\n=== Artist Lookup ====\n%s\n%s", req.URL, string(body))
	}
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response ArtistResultWrapper
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	result := response.Results[0]
	return result
}
