package handler

import "net/http"

func GetTimes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
