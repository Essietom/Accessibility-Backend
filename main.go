package Accessibility_Backend

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.Setup()
	route := mux.NewRouter()
	routes.RegisterRoutes(route)
	if err := http.ListenAndServe(":3000", route); err != nil {
		log.Fatal(err)
	}
}
