package middleware

import (
	"localhost/jwt"
	"localhost/models"
	"localhost/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var sessions = map[string]models.Session{}

func Login(c *gin.Context) {
	var inp models.LoginInput

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Json")
		return
	}
	if inp.Username == "" || inp.Password == "" {
		c.JSON(http.StatusBadRequest, "Bad username or password")
		return
	}

	user, flag := service.Login(inp.Username, inp.Password)

	if !flag {
		c.JSON(http.StatusBadRequest, "BAD REQUEST")
		return
	}

	jwt, err := jwt.GenerateJWT(*user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Error jwt generation")
		return
	}

	resp := models.LoginResponse{
		Token: jwt,
	}
	expirationTime := time.Now().Add(24 * time.Hour)

	sessions[jwt] = models.Session{
		Username: inp.Username,
		Expires:  expirationTime,
	}

	c.JSON(http.StatusCreated, resp)
	c.SetCookie(
		"token",
		jwt,
		expirationTime.Second(),
		"/",
		"localhost",
		false,
		true,
	)
}
