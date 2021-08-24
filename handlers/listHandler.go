package handlers

import (
	"go-assesment/controllers"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	controllers.ListController(w)
}
