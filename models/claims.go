package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
