package main

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/routes"
	"Accessibility-Backend/utilities"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	config, err := utilities.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

	database.Setup()
	route := mux.NewRouter()
	routes.RegisterRoutes(route)
	

	// port := os.Getenv("PORT")
	// if err := http.ListenAndServe(":" + port, route); err != nil {
	// 	log.Fatal(err)
	// }

	port := config.AppPort
	if err := http.ListenAndServe(":" + port, route); err != nil {
		log.Fatal(err)
	}

	
}
