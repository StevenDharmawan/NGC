package handlers

import (
	"encoding/json"
	"go-web-server/config"
	"go-web-server/models"
	"net/http"
)

func CreateHeroes(w http.ResponseWriter, r *http.Request) {
	var hero models.Heroes
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statement, err := config.DB.Prepare("INSERT INTO heroes(name, universe, skill, imageURL) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	result, err := statement.Exec(hero.Name, hero.Universe, hero.Skill, hero.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hero.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hero)
}

func CreateVillains(w http.ResponseWriter, r *http.Request) {
	var villain models.Villains
	err := json.NewDecoder(r.Body).Decode(&villain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statement, err := config.DB.Prepare("INSERT INTO villains(name, universe, imageURL) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()
	result, err := statement.Exec(villain.Name, villain.Universe, villain.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	villain.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villain)
}
