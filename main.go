package main

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/routes"
	"log"
	"net/http"
	// "os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	database.Setup()
	route := mux.NewRouter()
	routes.RegisterRoutes(route)
	if err := http.ListenAndServe(":3000", route); err != nil {
		log.Fatal(err)
	}

	// port := os.Getenv("PORT")
	// if err := http.ListenAndServe(":" + port, route); err != nil {
	// 	log.Fatal(err)
	// }
}
