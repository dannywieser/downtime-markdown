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
	sb.WriteString("\n")
	return sb.String()
}

func generateGenreMarkdown(book Book) string {
	var sb strings.Builder
	label := "Genre"
	if len(book.Genres) > 1 {
		label = "Genres"
	}
	sb.WriteString(utils.H2(label))
	for _, genre := range book.Genres {
		sb.WriteString(fmt.Sprintf("%s\n", utils.BuildTag("genres", genre)))
	}
	sb.WriteString("\n")
	return sb.String()
}

func generateDetailsMarkdown(book Book) string {
	var sb strings.Builder
	sb.WriteString(utils.H2("Details"))
	sb.WriteString(utils.KeyValue("Duration", "N/A"))
	sb.WriteString(utils.KeyValue("Status", utils.DefaultTag()))
	sb.WriteString(utils.KeyValue("Type", utils.BuildTag("books")))
	sb.WriteString(utils.KeyValue("Pages", fmt.Sprintf("%d", book.PageCount)))
	sb.WriteString(utils.KeyValue("Published", book.Published))
	sb.WriteString("\n")
	return sb.String()
}

func bookToMarkdown(book Book) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s\n", book.Title))
	sb.WriteString(generateAuthorMarkdown(book))
	sb.WriteString(generateDetailsMarkdown(book))
	sb.WriteString(generateGenreMarkdown(book))
	sb.WriteString(utils.H2("Synopsis"))
	sb.WriteString(book.Synopsis)
	sb.WriteString(fmt.Sprintf("\n- - - -\n ![](%s)\n", book.CoverImageUrl))

	return sb.String()
}

func GenerateBookMarkdown(title string) string {
	fmt.Printf("Generating markdown for \"%s\"\n", title)
	itunesResult := itunes.DoSearch(title, "ebook")
	olResult := openlibrary.DoSearch(title)
	book := mergeResults(itunesResult, olResult)
	return bookToMarkdown(book)
}
