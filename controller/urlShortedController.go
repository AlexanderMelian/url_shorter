package controller

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"localhost/jwt"
	"localhost/models"
	"localhost/service"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	hashedUrl := generateShortUrl(input.Url)
	uId, err := jwt.GetUserId(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Auth")
		return
	}
	service.SaveUrlShorted(hashedUrl, input.Url, uId)
	link := buildUrl(hashedUrl)
	c.JSON(http.StatusOK, link)
}

func DeleteUrlShorted(c *gin.Context) {
	url := c.Param("url")
	uId, err := jwt.GetUserId(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error deleting")
		return
	}
	er := service.DeleteUrlShorted(url, uId)
	if er != nil {
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
	shortUrl := base64.URLEncoding.EncodeToString(hash[:9])

	return shortUrl
}

func buildUrl(password string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	host := os.Getenv("HOST")

	var builder strings.Builder
	builder.Write([]byte(host))
	builder.Write([]byte("/"))
	builder.Write([]byte(password))
	return builder.String()
}
