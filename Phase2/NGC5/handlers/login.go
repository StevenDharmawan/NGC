package handlers

import (
	"Avengers/config"
	"Avengers/entity"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	var password string
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
		return
	}
	query := "SELECT * FROM users WHERE email = ?"
	err = config.DB.QueryRow(query, user.Email).Scan(&user.Id, &user.Email, &password, &user.FullName, &user.Age, &user.Occupation, &user.Role)
	if err != nil {
		panic(err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		panic(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Berhasil Login User")
}
