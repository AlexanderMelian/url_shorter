package controller

import (
	"localhost/database"
	"localhost/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectUrl(c *gin.Context) {
	url := c.Param("url")

	tinyUrl, err := database.FindUrlModelByShorted(url)
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	}

	c.Redirect(http.StatusPermanentRedirect, tinyUrl.UrlOriginal)
	service.UpdateLastUsedFromUrl(tinyUrl)

}
