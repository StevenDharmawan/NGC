package handlers

import (
	"NGC4/config"
	"NGC4/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := config.DB.Query("SELECT * FROM criminal_reports")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var criminalReports []models.CriminalReports
	for rows.Next() {
		var criminalReport models.CriminalReports

		err = rows.Scan(&criminalReport.CriminalId, &criminalReport.HeroId, &criminalReport.VillainId, &criminalReport.Description, &criminalReport.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		criminalReports = append(criminalReports, criminalReport)
	}
	response := models.Response{
		Code:    200,
		Message: "Berhasil Get data",
	}
	responseCriminalReports := models.ResponseCriminalReports{
		Response:        response,
		CriminalReports: criminalReports,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseCriminalReports)
}
