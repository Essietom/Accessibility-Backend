package main

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/routes"
	"Accessibility-Backend/utilities"
	"log"
	"net/http"
	"os"

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
	

	port := config.AppPort
	if port == ""{
		port  = os.Getenv("PORT")
	}
	if err := http.ListenAndServe(":" + port, route); err != nil {
		log.Fatal(err)
	}

	
}
