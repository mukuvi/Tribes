package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukuvi/Tribes/internal/data"
	"github.com/mukuvi/Tribes/internal/models"
	"github.com/mukuvi/Tribes/internal/utils"
)

// GET ALL
func GetTribes(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, data.Tribes)
}

// GET BY ID
func GetTribeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, tribe := range data.Tribes {
		if tribe.ID == id {
			utils.JSON(w, http.StatusOK, tribe)
			return
		}
	}

	utils.JSON(w, http.StatusNotFound, map[string]string{"error": "Tribe not found"})
}

// POST (CREATE)
func CreateTribe(w http.ResponseWriter, r *http.Request) {
	var tribe models.Tribe
	json.NewDecoder(r.Body).Decode(&tribe)

	data.Tribes = append(data.Tribes, tribe)

	utils.JSON(w, http.StatusCreated, tribe)
}

// PUT (UPDATE)
func UpdateTribe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, tribe := range data.Tribes {
		if tribe.ID == id {
			json.NewDecoder(r.Body).Decode(&data.Tribes[i])
			data.Tribes[i].ID = id
			utils.JSON(w, http.StatusOK, data.Tribes[i])
			return
		}
	}

	utils.JSON(w, http.StatusNotFound, map[string]string{"error": "Tribe not found"})
}

// DELETE
func DeleteTribe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, tribe := range data.Tribes {
		if tribe.ID == id {
			data.Tribes = append(data.Tribes[:i], data.Tribes[i+1:]...)
			utils.JSON(w, http.StatusOK, map[string]string{"message": "Deleted"})
			return
		}
	}

	utils.JSON(w, http.StatusNotFound, map[string]string{"error": "Tribe not found"})
}
