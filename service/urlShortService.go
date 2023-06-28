package service

import (
	"errors"
	"localhost/database"
	"localhost/models"
	"time"
)

func SaveUrlShorted(urlShorted string, url string, uId uint) error {
	tinyUrl := &models.TinyUrl{
		UrlShorted:  urlShorted,
		UrlOriginal: url,
		LastUsed:    time.Now(),
		UserId:      uId,
	}

	err := database.SaveUrlShorted(*tinyUrl)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUrlShorted(urlShorted string, uId uint) error {
	row, err := database.FindUrlModelByShorted(urlShorted)
	if err != nil {
		return errors.New("not found")
	}
	if row.UserId != uId {
		return errors.New("error non owner")
	}
	err = database.DeleteUrlShorted(row)
	if err != nil {
		return errors.New("error deleting url")
	}
	return nil
}

func FindUrlByShorted(shorted string) (string, error) {
	url, err := database.FindUrlByShorted(shorted)
	if err != nil {
		return "", err
	}
	return url, nil
}

func UpdateLastUsedFromUrl(urlModel models.TinyUrl) {
	urlModel.LastUsed = time.Now()
	database.UpdateLastUsedFromUrl(urlModel)
}
