package utils

import (
	"demo/models"
	"encoding/json"
	"net/http"
)

func ErrorMessage(w http.ResponseWriter, status int, message string) {
	byteData, _ := json.Marshal(models.Error{
		Status:  status,
		Message: message,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(byteData)
}
