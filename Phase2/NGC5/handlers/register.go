package handlers

import (
	"Avengers/config"
	"Avengers/entity"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		panic(err)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err)
		return
	}
	query := "INSERT INTO users(email, password, fullname, age, occupation, role) VALUES(?, ?, ?, ?, ?, 'Admin')"
	_, err = config.DB.Exec(query, user.Email, hashPassword, user.FullName, user.Age, user.Occupation)
	if err != nil {
		panic(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Berhasil Register User")
}
