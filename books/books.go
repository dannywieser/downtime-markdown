package books

import (
	"dgw/downtime/itunes"
	"dgw/downtime/openlibrary"
	"dgw/downtime/utils"
	"fmt"
	"strings"
)

func mergeResults(itunes itunes.SearchResult, ol openlibrary.OpenLibraryBook) Book {
	var book Book

	// info from open library
	book.Title = ol.Title
	book.Published = ol.Published
	book.PageCount = ol.Pages
	book.CoverImageUrl = ol.Cover.Medium
	for _, olAuthor := range ol.Authors {
		book.Authors = append(book.Authors, Author{
			Name: olAuthor.Name,
			Url:  olAuthor.Url,
		})
	}

	// info from itunes
	book.Genres = itunes.Genres
	book.Synopsis = itunes.Synopsis

	return book
}

func generateAuthorMarkdown(book Book) string {
	var sb strings.Builder
	label := "Author"
	if len(book.Authors) > 1 {
		label = "Authors"
	}
	sb.WriteString(utils.H2(label))
	for _, author := range book.Authors {
		sb.WriteString(fmt.Sprintf("%s\n", utils.BuildTag("authors", author.Name)))
	}
	return sb.String()
}

func bookToMarkdown(book Book) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s\n", book.Title))

	sb.WriteString(generateAuthorMarkdown(book))
	sb.WriteString(utils.H2("Details"))
	sb.WriteString(utils.KeyValue("Pages", fmt.Sprintf("%d", book.PageCount)))
	sb.WriteString(utils.KeyValue("Published", book.Published))

	return sb.String()

}

func GenerateBookMarkdown(title string) string {
	fmt.Printf("Generating markdown for \"%s\"\n", title)
	itunesResult := itunes.DoSearch(title, "ebook")
	olResult := openlibrary.DoSearch(title)
	book := mergeResults(itunesResult, olResult)
	return bookToMarkdown(book)
}
