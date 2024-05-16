package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res := map[string]any{
		"Error": false,
		"Data": todos,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
