package openlibrary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	booksPath = "https://openlibrary.org/api/books"
)

func getISBNKey(isbn string) string {
	return fmt.Sprintf("ISBN:%s", isbn)
}

func getBook(isbn string) book {
	client := http.Client{Timeout: timeout}

	req, err := http.NewRequest(http.MethodGet, booksPath, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("bibkeys", getISBNKey(isbn))
	q.Add("jscmd", "data")
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL)

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

	jsonMap := make(map[string]book)
	jsonErr := json.Unmarshal(body, &jsonMap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	book := jsonMap[getISBNKey(isbn)]
	return book
}
