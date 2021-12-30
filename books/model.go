package books

type Book struct {
	Title         string
	Authors       []Author
	Genres        []string
	CoverImageUrl string
	PageCount     int
	Published     string
	Synopsis      string
}

type Author struct {
	Name string
	Url  string
}
