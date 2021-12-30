package openlibrary

import (
	"dgw/downtime/utils"
	"encoding/json"
	"fmt"
	"log"
)

const (
	booksPath = "https://openlibrary.org/api/books"
)

func getISBNKey(isbn string) string {
	return fmt.Sprintf("ISBN:%s", isbn)
}

func buildParams(isbn string) map[string]string {
	params := make(map[string]string)
	params["bibkeys"] = getISBNKey(isbn)
	params["jscmd"] = "data"
	params["format"] = "json"
	return params
}

func getBook(isbn string) OpenLibraryBook {
	body := utils.DoGet(booksPath, buildParams(isbn))

	isbnMap := make(map[string]OpenLibraryBook)
	jsonErr := json.Unmarshal(body, &isbnMap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return isbnMap[getISBNKey(isbn)]
}
