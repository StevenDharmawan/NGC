package main

import (
	"fmt"
	"go-web-server/config"
	"go-web-server/handlers"
	"net/http"
)

func main() {
	config.InitDB("root:@tcp(localhost:3306)/Avengers")
	defer config.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/heroes", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHeroes(w, r)
	})
	mux.HandleFunc("/villains", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetVillains(w, r)
	})
	mux.HandleFunc("/create/heroes", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateHeroes(w, r)
	})
	mux.HandleFunc("/create/villains", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateVillains(w, r)
	})

	fmt.Println("Running server on port :8080")

	// running web server on local env
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
