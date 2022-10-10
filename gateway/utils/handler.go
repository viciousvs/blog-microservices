package utils

import "net/http"

const ContentType = "Content-Type"
const ApplJson = "application/json"

func ErrorHandler(w http.ResponseWriter, err string, status int) {
	http.Error(w, err, status)
}
