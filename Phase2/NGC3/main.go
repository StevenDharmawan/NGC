package main

import (
	"NGC3/config"
	"NGC3/handlers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	config.InitDB("root:@tcp(localhost:3306)/Avengers")
	defer config.DB.Close()

	router := httprouter.New()
	router.GET("/inventories", handlers.GetAll)
	router.GET("/inventories/:id", handlers.GetByID)
	router.POST("/inventories", handlers.Create)
	router.PUT("/inventories/:id", handlers.UpdateInventory)
	router.DELETE("/inventories/:id", handlers.Delete)
	fmt.Println("Running server on port:8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}
