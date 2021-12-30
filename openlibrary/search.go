package openlibrary

import (
	"dgw/downtime/utils"
	"encoding/json"
	"fmt"
	"log"
)

const (
	searchPath = "http://openlibrary.org/search.json"
)

func DoSearch(title string) OpenLibraryBook {
	params := make(map[string]string)
	params["q"] = fmt.Sprintf("title:%s", title)
	params["fields"] = "title,author_name,isbn"
	responseBody := utils.DoGet(searchPath, params)

	var response searchResultWrapper
	jsonErr := json.Unmarshal(responseBody, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var result OpenLibraryBook
	if response.NumFound > 0 {
		fmt.Printf("  ✅ Open Library API\n")
		result = getBook(response.Docs[0].ISBN[0])
	} else {
		fmt.Printf("  ❌ Open Library API\n")
	}

	return result
}
