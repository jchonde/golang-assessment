package handlers

import (
	"go-assesment/controllers"
	"net/http"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	word := r.URL.Query().Get("word")

	controllers.RemoveController(word)
}
