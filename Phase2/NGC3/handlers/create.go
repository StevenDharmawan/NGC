package handlers

import (
	"NGC3/config"
	"NGC3/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var inventory models.Inventories

	// decode body request then assign to models.Book
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement, err := config.DB.Prepare("INSERT INTO inventories(name, description, stock, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	result, err := statement.Exec(inventory.Name, inventory.Description, inventory.Stock, inventory.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	inventory.ItemCode = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}
