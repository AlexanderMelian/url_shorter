package middleware

import (
	"localhost/jwt"
	"localhost/models"
	"localhost/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.GetHeader("Authorization")

		if cookie == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		err := jwt.ValidateAuthToken(cookie)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

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

	user, err := service.Login(inp.Username, inp.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, "BAD REQUEST")
		return
	}

	jwt, err := jwt.GenerateJWT(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Error jwt generation")
		return
	}

	resp := models.LoginResponse{
		Token: jwt,
	}
	expirationTime := time.Now().Add(24 * time.Hour)

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
