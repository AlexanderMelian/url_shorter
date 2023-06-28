package jwt

import (
	"errors"
	"fmt"
	"localhost/models"
	"localhost/service"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey []byte

func init() {
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if secretKey == nil {
		panic("Could not load JWT secret key")
	}
}

func GenerateJWT(u models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["username"] = u.Username
	claims["email"] = u.Email
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

// only return if token is not valid
func ValidateAuthToken(signedToken string) error {
	claims := &jwt.MapClaims{}
	splitedToken := strings.Split(signedToken, "Bearer ")

	token, er := jwt.ParseWithClaims(
		splitedToken[1],
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)

	if er != nil {
		return errors.New("error generating token")
	}
	if !token.Valid {
		return errors.New("TOKEN NOT VALID")
	}
	return nil
}

func GetUserId(signedToken string) (uint, error) {
	signedToken = strings.Replace(signedToken, "Bearer ", "", 1)
	var name string
	token, _, err := new(jwt.Parser).ParseUnverified(signedToken, jwt.MapClaims{})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		name = fmt.Sprint(claims["username"])
	}
	if name == "" {
		return 0, errors.New("invalid token payload")
	}
	uId, err := service.FindByUsername(name)
	if err != nil {
		return 0, err
	}
	return uId.ID, nil
}
