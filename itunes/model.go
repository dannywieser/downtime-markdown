package itunes

import "time"

type SearchResultWrapper struct {
	ResultCount int            `json:"resultCount"`
	Results     []SearchResult `json:"results"`
}

type SearchResult struct {
	TrackName       string `json:"trackName"`
	LongDescription string `json:"longDescription"`
	Description     string `json:"description"`
	ArtistName      string `json:"ArtistName"`
	ArtistIds       []int  `json:"ArtistIds"`
	ArtworkUrl      string `json:"artworkUrl100"`
	Genres          []string
	ReleaseDate     time.Time
	Artists         []string
	MediaType       string
}

type ArtistResultWrapper struct {
	ResultCount int            `json:"resultCount"`
	Results     []ArtistResult `json:"results"`
}

type ArtistResult struct {
	Name string `json:"artistName"`
	Url  string `json:"artistLinkUrl"`
}

type SearchParams struct {
	Title     string
	MediaType string
	Debug     bool
}
