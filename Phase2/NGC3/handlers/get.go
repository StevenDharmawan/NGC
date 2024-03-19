package handlers

import (
	"NGC3/config"
	"NGC3/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := config.DB.Query("SELECT * FROM inventories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var inventories []models.Inventories
	for rows.Next() {
		var inventory models.Inventories

		err = rows.Scan(&inventory.ItemCode, &inventory.Name, &inventory.Description, &inventory.Stock, &inventory.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		inventories = append(inventories, inventory)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventories)
}

func GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	var inventory models.Inventories

	err := config.DB.QueryRow("SELECT * FROM inventories WHERE item_code = ?", id).Scan(&inventory.ItemCode, &inventory.Name, &inventory.Description, &inventory.Stock, &inventory.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}
