package models

import "time"

type TinyUrl struct {
	ID          uint      `json:"id"`
	UrlShorted  string    `json:"url-shorted"`
	UrlOriginal string    `json:"url-original"`
	UserId      uint      `json:"created-by"`
	LastUsed    time.Time `json:"last-used"`
}
