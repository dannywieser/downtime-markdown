package openlibrary

type searchResultWrapper struct {
	NumFound int
	Docs     []OpenLibrarySearchResult
}

type OpenLibrarySearchResult struct {
	Seed             []string
	Title            string
	Author           []string `json:"author_name"`
	ISBN             []string `json:"isbn"`
	CoverId          int      `json:"cover_i"`
	FirstPublishYear int      `json:"first_publish_year"`
	Pages            int      `json:"number_of_pages_median"`
	CoverImageUrl    string
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
