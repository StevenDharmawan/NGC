package handlers

import (
	"NGC4/config"
	"NGC4/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var criminalReport models.CriminalReports

	err := json.NewDecoder(r.Body).Decode(&criminalReport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement, err := config.DB.Prepare("INSERT INTO criminal_reports(hero_id, villain_id, description, date) VALUES (?, ?, ?, CURDATE())")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	result, err := statement.Exec(criminalReport.HeroId, criminalReport.VillainId, criminalReport.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	criminalReport.CriminalId = id
	response := models.Response{
		Code:    200,
		Message: "Berhasil membuat data",
	}
	responseCriminalReport := models.ResponseCriminalReport{
		Response:       response,
		CriminalReport: criminalReport,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseCriminalReport)
}
