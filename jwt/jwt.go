package jwt

import (
	"errors"
	"localhost/models"
	"strings"

	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("mySuperSecretKEY123")

func GenerateJWT(u models.User) (string, error) {
	payload := jwt.MapClaims{
		"username": u.Username,
		"email":    u.Email,
		"expires":  time.Now().UTC().Add(time.Hour * 24),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

// only return if token is not valid
func ValidateAuthToken(signedToken string) error {
	claims := &models.Claim{}

	splitedToken := strings.Split(signedToken, "Bearer ")

	token, er := jwt.ParseWithClaims(
		splitedToken[1],
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if er != nil {
		return er
	}

	if !token.Valid {
		return errors.New("TOKEN NOT VALID")
	}
	return nil
}
