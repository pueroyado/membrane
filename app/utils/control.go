package utils

import (
	"net/http"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xApiKey := r.Header.Get("x-api-key")
		if xApiKey == "" {
			ErrorMessage(w, http.StatusUnauthorized, "Invalid x-api-key")
			return
		}

		next.ServeHTTP(w, r)
	})
}
