package jwt

import "github.com/golang-jwt/jwt"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
