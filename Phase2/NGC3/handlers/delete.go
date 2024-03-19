package handlers

import (
	"NGC3/config"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	_, err := config.DB.Exec("DELETE FROM inventories WHERE item_code = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Inventory deleted successfully")
}
