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
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("couldn't load")
	}
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
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
		return er
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
	uId, find := service.FindByUsername(name)
	if !find {
		return 0, errors.New("Invalid username")
	}
	return uId.ID, nil
}
