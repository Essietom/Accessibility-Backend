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
	}).Methods("GET")
	router.HandleFunc("/criteria/", controllers.CreateCriteria).Methods("POST")
	router.HandleFunc("/criteria/", controllers.GetCriteria).Methods("GET")
	router.HandleFunc("/criteria/{criteriaId}", controllers.GetCriteriaById).Methods("GET")
	router.HandleFunc("/criteria/{criteriaId}", controllers.UpdateCriteria).Methods("PUT")
	router.HandleFunc("/criteria/{criteriaId}", controllers.DeleteCriteria).Methods("DELETE")

	router.HandleFunc("/webpage/", controllers.SaveWebpageScans).Methods("POST")
	router.HandleFunc("/webpage/", controllers.GetWebpageScan).Methods("GET")
	router.HandleFunc("/webpage/{webpageId}", controllers.GetWebpageScanById).Methods("GET")
	router.HandleFunc("/webpage/field/{webpageField}", controllers.GetWebpageByField).Methods("GET")
	router.HandleFunc("/webpage/{webpageId}", controllers.UpdateWebpageScan).Methods("PUT")
	router.HandleFunc("/webpage/{webpageId}", controllers.DeleteWebpageScan).Methods("DELETE")

	router.HandleFunc("/issue/{webpageId}", controllers.AddIssue).Methods("POST")
	router.HandleFunc("/issue/{webpageId}", controllers.GetAllIssuesforWebpageId).Methods("GET")
	router.HandleFunc("/issue/{webpageId}/{issueId}", controllers.UpdateIssueByIssueIdAndWebpageId).Methods("PUT")
	router.HandleFunc("/issue/{webpageId}/{issueId}", controllers.DeleteIssueByIssueIdAndWebpageId).Methods("DELETE")

	router.HandleFunc("/website/", controllers.CreateWebsite).Methods("POST")
	router.HandleFunc("/website/", controllers.GetAllWebsites).Methods("GET")
	router.HandleFunc("/website/{websiteId}", controllers.GetWebsiteById).Methods("GET")
	router.HandleFunc("/website/{websiteId}", controllers.UpdateWebsite).Methods("PUT")

}
