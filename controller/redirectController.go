package controller

import (
	"localhost/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectUrl(c *gin.Context) {
	url := c.Param("url")

	tinyUrl, find := service.FindUrlByShorted(url)
	if !find {
		c.JSON(http.StatusNotFound, "Not Found")
	}

	c.Redirect(http.StatusPermanentRedirect, tinyUrl.UrlOriginal)
	service.UpdateLastUsedFromUrl(*tinyUrl)

}
