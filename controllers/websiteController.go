package controllers

import (
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"net/http"

	"github.com/gorilla/mux"
)

var NewWebsite entity.Website

func CreateWebsite(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	websiteRequest := &dto.WebsiteRequestBody{}
	utilities.ParseBody(r, websiteRequest)
	website, err := model.CreateWebsite(websiteRequest)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(website, w, r)
}

func GetAllWebsites(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	websites, err := model.GetAllWebsites()

	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(websites, w, r)
}

func GetWebsiteById(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	vars := mux.Vars(r)
	websiteId := vars["websiteId"]
	websiteDetails, err := model.GetWebsiteById(websiteId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(websiteDetails, w, r)
}

func UpdateWebsite(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	var updateWebsite = &entity.Website{}
	utilities.ParseBody(r, updateWebsite)
	vars := mux.Vars(r)
	websiteId := vars["websiteId"]

	websiteDetails, err := model.GetWebsiteById(websiteId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}

	if updateWebsite.Name != "" {
		websiteDetails.Name = updateWebsite.Name
	}
	if updateWebsite.Url != "" {
		websiteDetails.Url = updateWebsite.Url
	}
	model.UpdateWebsite(websiteDetails, websiteId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(websiteDetails, w, r)

}
