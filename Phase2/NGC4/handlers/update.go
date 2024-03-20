package handlers

import (
	"NGC4/config"
	"NGC4/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var criminalReport models.CriminalReports

	err := json.NewDecoder(r.Body).Decode(&criminalReport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ps.ByName("id")
	_, err = config.DB.Exec("UPDATE criminal_reports SET hero_id = ?, villain_id = ?, description = ? WHERE criminal_id = ?", criminalReport.HeroId, criminalReport.VillainId, criminalReport.Description, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := models.Response{
		Code:    200,
		Message: "Berhasil Update Data",
	}
	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
