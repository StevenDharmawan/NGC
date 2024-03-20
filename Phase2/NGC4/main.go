package main

import (
	"NGC4/config"
	"NGC4/handlers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	config.InitDB("root:@tcp(localhost:3306)/Avengers?parseTime=true")
	defer config.DB.Close()

	router := httprouter.New()
	router.GET("/criminalreport", handlers.GetAll)
	router.POST("/criminalreport", handlers.Create)
	router.PUT("/criminalreport/:id", handlers.Update)
	router.DELETE("/criminalreport/:id", handlers.Delete)

	fmt.Println("Running server on port:8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server : ", err.Error())
	}
}
