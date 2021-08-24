package handlers

import (
	"go-assesment/controllers"
	"io"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	word := r.URL.Query().Get("word")

	wordIsAdded, err := controllers.AddController(word)

	if wordIsAdded && err == nil {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Word added \n")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Something went wrong \n")
	}

}
