package middleware

import (
	"localhost/jwt"
	"net/http"

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
