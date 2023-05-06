package service

import (
	"errors"
	"fmt"
	"localhost/models"
	"localhost/setup"
	"time"

	"gorm.io/gorm"
)

func SaveUrlShorted(urlShorted string, url string, uId uint) {
	db := setup.DB
	tinyUrl := &models.TinyUrl{
		UrlShorted:  urlShorted,
		UrlOriginal: url,
		LastUsed:    time.Now(),
		UserId:      uId,
	}

	result := db.Create(tinyUrl)

	if result.Error != nil {
		panic(result.Error)
	}
}

func DeleteUrlShorted(urlShorted string, uId uint) error {
	db := setup.DB
	row, exist := FindUrlByShorted(urlShorted)
	if !exist {
		return errors.New("not found")
	}
	if row.UserId != uId {
		return errors.New("Error now owner")
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

func UpdateLastUsedFromUrl(urlModel models.TinyUrl) {
	db := setup.DB
	urlModel.LastUsed = time.Now()
	fmt.Println(urlModel)
	db.Save(&urlModel)
}
