package database

import (
	"errors"
	"localhost/models"
)

func SaveUrlShorted(tinyUrl models.TinyUrl) error {
	err := DB.Create(&tinyUrl).Error

	if err != nil {
		return errors.New("error creating tinyurl")
	}
	return nil
}

func UpdateLastUsedFromUrl(tinyUrl models.TinyUrl) error {
	err := DB.Save(&tinyUrl)
	if err != nil {
		return errors.New("error updating last_used")
	}
	return nil
}

func DeleteUrlShorted(urlShorted models.TinyUrl) error {
	err := DB.Where("id = ?", urlShorted.ID).Delete(&urlShorted).Error
	if err != nil {
		return errors.New("error deleting url")
	}
	return nil
}

func FindUrlByShorted(urlShorted string) (string, error) {
	var sh models.TinyUrl
	println("ASDASDASD")
	err := DB.Where("url_shorted = ?", urlShorted).First(&sh).Error
	if err != nil {
		return "", errors.New("not found")
	}
	return sh.UrlOriginal, nil
}
func FindUrlModelByShorted(urlShorted string) (models.TinyUrl, error) {
	var sh models.TinyUrl
	err := DB.Where("url_shorted = ?", urlShorted).First(&sh).Error
	if err != nil {
		return models.TinyUrl{}, errors.New("not found")
	}
	return sh, nil
}
