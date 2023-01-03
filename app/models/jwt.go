package models

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwtSecretKey = []byte("wJH8fsVxa8nnWTE5W14a")

type JwtPayload struct {
	Token    string `json:"token"`
	TokenExp int64  `json:"tokenExp"`
}

func CreateJwt(userId int32) *JwtPayload {
	expTime := time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    expTime,
		"userId": userId,
	})

	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString(JwtSecretKey)
	jwtPayload := &JwtPayload{
		Token:    tokenString,
		TokenExp: expTime,
	}

	return jwtPayload
}
