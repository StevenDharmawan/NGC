package handlers

import (
	"encoding/json"
	"go-web-server/config"
	"go-web-server/models"
	"net/http"
)

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM heroes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []models.Heroes
	for rows.Next() {
		var hero models.Heroes

		err = rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		heroes = append(heroes, hero)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func GetVillains(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM villains")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var villains []models.Villains
	for rows.Next() {
		var villain models.Villains

		err = rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		villains = append(villains, villain)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villains)
}
