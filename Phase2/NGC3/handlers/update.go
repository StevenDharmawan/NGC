package handlers

import (
	"NGC3/config"
	"NGC3/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func UpdateInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var inventory models.Inventories

	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ps.ByName("id")
	_, err = config.DB.Exec("UPDATE inventories SET name = ?, description = ?, stock = ?, status = ? WHERE item_code = ?", inventory.Name, inventory.Description, inventory.Stock, inventory.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Inventory updated successfully")
}
