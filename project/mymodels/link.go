package mymodels

type NewLink struct {
	URL          string `json:"url"`
}

type NewLinkResponse struct {
	ShortURL          string `json:"shorturl"`
}