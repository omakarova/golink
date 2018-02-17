package mymodels

type NewLink struct {
	URL          string `json:"url"`
}

type NewLinkResponse struct {
	ShortURL          string `json:"shorturl"`
}

type LinkInfo struct {
	ShortURL          string `json:"shorturl"`
	LongURL          string `json:"longurl"`
	NumberOfClicks	int `json:"numberofclicks"`
}



