package handlers

import (
	"go-assesment/controllers"
	"io"
	"net/http"
)

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	word1 := r.URL.Query().Get("word1")
	word2 := r.URL.Query().Get("word2")

	solution, err := controllers.CompareController(word1, word2)

	if solution != "" || err == nil {
		io.WriteString(w, solution)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
