package controllers

import (
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var NewWebsite entity.Website

func CreateWebsite(w http.ResponseWriter, r *http.Request) {
	vi := &entity.Website{}
	utilities.ParseBody(r, vi)
	v, _ := model.CreateWebsite(vi)
	res, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllWebsites(w http.ResponseWriter, r *http.Request) {
	criteria, err := model.GetAllWebsites()
	if err != nil {

	}
	res, _ := json.Marshal(criteria)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetWebsiteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	websiteId := vars["websiteId"]
	websiteDetails, err := model.GetWebsiteById(websiteId)
	if err != nil {

	}
	res, _ := json.Marshal(websiteDetails)
	w.Header().Set("Content-Type", "pkglicatio/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateWebsite(w http.ResponseWriter, r *http.Request) {
	var updateWebsite = &entity.Website{}
	utilities.ParseBody(r, updateWebsite)
	vars := mux.Vars(r)
	websiteId := vars["websiteId"]

	websiteDetails, err := model.GetWebsiteById(websiteId)
	if err != nil {
		println("tomiiiii", err)
		return

	}

	if updateWebsite.Name != "" {
		websiteDetails.Name = updateWebsite.Name
	}
	if updateWebsite.Url != "" {
		websiteDetails.Url = updateWebsite.Url
	}
	model.UpdateWebsite(websiteDetails, websiteId)
	res, _ := json.Marshal(websiteDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
