package handlers

import (
	"go-assesment/controllers"
	"net/http"
)

func FindLongestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	controllers.FindLongestController(w)
}
