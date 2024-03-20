package models

import "time"

type CriminalReports struct {
	CriminalId  int64     `json:"criminal_id"`
	HeroId      int       `json:"hero_id"`
	VillainId   int       `json:"villain_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Response struct {
	Code    int    `json:"response"`
	Message string `json:"message"`
}

type ResponseCriminalReport struct {
	Response       Response
	CriminalReport CriminalReports `json:"criminal_report"`
}

type ResponseCriminalReports struct {
	Response        Response
	CriminalReports []CriminalReports `json:"criminal_reports"`
}
