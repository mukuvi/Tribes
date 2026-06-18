package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukuvi/Tribes/internal/handlers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := `
			Welcome to Kenya Tribes API

			Available Endpoints:

			GET all tribes
			GET /tribes

			GET tribe by ID
			GET /tribes/1

			CREATE tribe
			POST /tribes

			UPDATE tribe
			PUT /tribes/1

			DELETE tribe
			DELETE /tribes/1
			`
		w.Write([]byte(message))
	}).Methods("GET")

	r.HandleFunc("/tribes", handlers.GetTribes).Methods("GET")
	r.HandleFunc("/tribes/{id}", handlers.GetTribeByID).Methods("GET")

	r.HandleFunc("/tribes", handlers.CreateTribe).Methods("POST")
	r.HandleFunc("/tribes/{id}", handlers.UpdateTribe).Methods("PUT")
	r.HandleFunc("/tribes/{id}", handlers.DeleteTribe).Methods("DELETE")

	return r
}
