package handlers

import (
	"NGC4/config"
	"NGC4/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	_, err := config.DB.Exec("DELETE FROM criminal_reports WHERE criminal_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := models.Response{
		Code:    200,
		Message: "Berhasil Delete Data",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
