package main

import (
	"Avengers/config"
	"Avengers/handlers"
	"net/http"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", handlers.Register)
	mux.HandleFunc("POST /login", handlers.Login)
	http.ListenAndServe(":8080", mux)

}
