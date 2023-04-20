package service

import (
	"localhost/models"
	"localhost/setup"
	"time"
)

func SaveUrlShorted(urlShorted string, url string) {
	db := setup.DB
	tinyUrl := &models.TinyUrl{
		UrlShorted:  urlShorted,
		UrlOriginal: url,
		LastUsed:    time.Now(),
		UserId:      0,
	}

	result := db.Create(tinyUrl)

	if result.Error != nil {
		panic(result.Error)
	}
}
