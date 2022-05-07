package routes

import (
	"Accessibility-Backend/controllers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		json.NewEncoder(rw).Encode(map[string]string{"data": "Hello from Mux & mongoDB"})
	}).Methods("GET", "OPTIONS")
	router.HandleFunc("/criteria", controllers.CreateCriteria).Methods("POST", "OPTIONS")
	router.HandleFunc("/criteria", controllers.GetCriteria).Methods("GET", "OPTIONS")
	router.HandleFunc("/criteria/{criteriaId}", controllers.GetCriteriaById).Methods("GET", "OPTIONS")
	router.HandleFunc("/criteria/{criteriaId}", controllers.UpdateCriteria).Methods("PUT", "OPTIONS")
	router.HandleFunc("/criteria/{criteriaId}", controllers.DeleteCriteria).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/webpage", controllers.SaveWebpageScans).Methods("POST", "OPTIONS")
	//router.HandleFunc("/webpage", controllers.GetWebpageScan).Methods("GET")
	router.HandleFunc("/webpage/{webpageId}", controllers.GetWebpageScanById).Methods("GET", "OPTIONS")
	router.HandleFunc("/webpage", controllers.GetWebpageByField).Methods("GET", "OPTIONS")
	//router.HandleFunc("/webpage/{webpageId}", controllers.UpdateWebpageScan).Methods("PUT")
	router.HandleFunc("/webpage/{webpageId}", controllers.DeleteWebpageScan).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/issue/{webpageId}", controllers.AddIssue).Methods("POST", "OPTIONS")
	router.HandleFunc("/issue/{webpageId}", controllers.GetAllIssuesforWebpageId).Methods("GET", "OPTIONS")
	router.HandleFunc("/issue", controllers.UpdateIssueByIssueIdAndWebpageId).Methods("PUT", "OPTIONS")
	router.HandleFunc("/issue", controllers.DeleteIssueByIssueIdAndWebpageId).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/occurence", controllers.DeleteOccurenceIdAndIssueIdAndWebpageId).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/website", controllers.CreateWebsite).Methods("POST", "OPTIONS")
	router.HandleFunc("/website", controllers.GetAllWebsites).Methods("GET", "OPTIONS")
	router.HandleFunc("/website/{websiteId}", controllers.GetWebsiteById).Methods("GET", "OPTIONS")
	router.HandleFunc("/website/{websiteId}", controllers.UpdateWebsite).Methods("PUT", "OPTIONS")

}
