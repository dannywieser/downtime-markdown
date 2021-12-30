package utils

import (
	"dgw/downtime/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const timeout = time.Second * 20

func DoGet(url string, params map[string]string) []byte {
	client := http.Client{Timeout: timeout}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	if config.DebugMode {
		fmt.Printf("URL: %s\n Result %s\n", req.URL, string(body))
	}

	return body
}
