package routes

import (
	"net/http"

	"github.com/mukuvi/Tribes/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Kenya Tribes API "))
	}).Methods("GET")

	r.HandleFunc("/tribes", handlers.GetTribes).Methods("GET")

	return r
}
