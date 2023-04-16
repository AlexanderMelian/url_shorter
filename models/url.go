package models

type TinyUrl struct {
	ID          uint   `json:"id"`
	UrlShorted  string `json:"url-shorted"`
	UrlOriginal string `json:"url-original"`
	CreatedBy   uint   `json:"created-by"`
}
