package itunes

import "time"

const apiPath = "https://itunes.apple.com/search"

type SearchResultWrapper struct {
	ResultCount int            `json:"resultCount"`
	Results     []SearchResult `json:"results"`
}

type SearchResult struct {
	TrackName       string `json:"trackName"`
	LongDescription string `json:"longDescription"`
	Description     string `json:"description"`
	ArtistName      string `json:"ArtistName"`
	ArtworkUrl      string `json:"artworkUrl100"`
	Genres          []string
	ReleaseDate     time.Time
}

type SearchParams struct {
	Title     string
	MediaType string
	Debug     bool
}
