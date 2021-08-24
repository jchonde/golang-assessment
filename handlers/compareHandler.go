package handlers

import (
	"go-assesment/controllers"
	"net/http"
)

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	word1 := r.URL.Query().Get("word1")
	word2 := r.URL.Query().Get("word2")

	controllers.CompareController(w, word1, word2)
}
