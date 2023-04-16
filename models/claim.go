package models

import "github.com/golang-jwt/jwt"

type Claim struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
