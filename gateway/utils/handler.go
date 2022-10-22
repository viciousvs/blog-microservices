package utils

import (
	"encoding/json"
	"net/http"
)

const ContentType = "Content-Type"
const ApplJson = "application/json"

func ErrorHandler(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(map[string]string{"message": msg})
	if err != nil {
		return
	}
}
