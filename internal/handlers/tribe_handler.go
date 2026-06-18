package handlers

import (
	"net/http"

	"github.com/mukuvi/Tribes/internal/data"
	"github.com/mukuvi/Tribes/internal/utils"
)

func GetTribes(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, data.Tribes)
}
