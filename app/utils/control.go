package utils

import (
	"demo/models"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// x-api-key
		xApiKey := r.Header.Get("x-api-key")
		if xApiKey == "" {
			ErrorMessage(w, http.StatusUnauthorized, "Invalid x-api-key")
			return
		}

		// Authorization
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			ErrorMessage(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		splitToken := strings.Split(reqToken, "Bearer ")
		token, errParse := jwt.Parse(splitToken[1], func(t *jwt.Token) (interface{}, error) {
			return models.JwtSecretKey, nil
		})

		if errors.Is(errParse, jwt.ErrTokenMalformed) || token.Valid == false {
			ErrorMessage(w, http.StatusUnauthorized, "Token invalid")
			return
		} else if errors.Is(errParse, jwt.ErrTokenExpired) {
			ErrorMessage(w, http.StatusUnauthorized, "Token expired")
			return
		} else if errors.Is(errParse, jwt.ErrTokenNotValidYet) {
			ErrorMessage(w, http.StatusUnauthorized, "Token not valid yet")
			return
		}

		next.ServeHTTP(w, r)
	})
}
