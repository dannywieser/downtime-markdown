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

type bookWrapper struct {
	Books map[string]book
}

type book struct {
	Title     string
	Published string `json:"publish_date"`
	Pages     int    `json:"number_of_pages"`
	Cover     cover
}

type cover struct {
	Small  string
	Medium string
	Large  string
}
