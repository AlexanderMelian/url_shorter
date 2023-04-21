package service

import (
	"errors"
	"localhost/models"
	"localhost/setup"
	"time"

	"gorm.io/gorm"
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

func DeleteUrlShorted(urlShorted string) error {
	db := setup.DB
	row, exist := FindUrlByShorted(urlShorted)
	if !exist {
		return errors.New("not found")
	}
	result := db.Where("id = ?", row.ID).Delete(&row)
	if result.Error != nil {
		return errors.New("error deleting url")
	}
	return nil
}

func FindUrlByShorted(shorted string) (*models.TinyUrl, bool) {
	db := setup.DB
	var sh models.TinyUrl
	result := db.Where("url_shorted = ?", shorted).First(&sh)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false
	}
	return &sh, true

}
