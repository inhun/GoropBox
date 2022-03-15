package models

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	NickName string `json:"NickName"`
}
