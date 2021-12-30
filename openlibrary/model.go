package openlibrary

type searchResultWrapper struct {
	NumFound int
	Docs     []docsResult
}

type docsResult struct {
	Seed   []string
	Title  string
	Author string
	ISBN   []string `json:"isbn"`
}

type OpenLibraryBook struct {
	Title     string
	Published string `json:"publish_date"`
	Pages     int    `json:"number_of_pages"`
	Cover     OpenLibraryCover
	Authors   []OpenLibraryAuthor
}

type OpenLibraryAuthor struct {
	Url  string
	Name string
}

type OpenLibraryCover struct {
	Small  string
	Medium string
	Large  string
}
