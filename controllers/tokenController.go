package controllers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func generateTokenPair(id uint, name string, role string) (map[string]string, error) {

	type jwtCustomClaims struct {
		Sub  uint   `json:"sub"`
		Name string `json:"name"`
		Role string `json:"role"`
		jwt.StandardClaims
	}

	claims := &jwtCustomClaims{
		id,
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 16).Unix(),
		},
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}
