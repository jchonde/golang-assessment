package handlers

import (
	"go-assesment/controllers"
	"io"
	"net/http"
	"reflect"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	word := r.URL.Query().Get("word")

	solution, err := controllers.RemoveController(word)

	if !reflect.DeepEqual(solution, "error") && err == nil {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, solution+"\n")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, solution+"\n")
	}

}
