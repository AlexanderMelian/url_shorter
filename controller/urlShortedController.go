package controller

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"localhost/models"
	"localhost/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUrlShorted(c *gin.Context) {
	var input models.UrlInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Json")
		return
	}
	if input.Url == "" {
		c.JSON(http.StatusBadRequest, "Input error")
		return
	}

	hashedPassword := generateShortUrl(input.Url)
	service.SaveUrlShorted(hashedPassword, input.Url)
}

func DeleteUrlShorted(c *gin.Context) {
	url := c.Param("url")
	err := service.DeleteUrlShorted(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error deleting")
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func generateShortUrl(url string) string {
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256([]byte(url + string(randomBytes)))

	shortUrl := base64.URLEncoding.EncodeToString(hash[:8])

	return shortUrl
}
