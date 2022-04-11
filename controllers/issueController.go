package controllers

import (
	"Accessibility-Backend/entity"
	"Accessibility-Backend/models"
	"Accessibility-Backend/utilities"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var Issue entity.Issue

var validate *validator.Validate

func AddIssue(w http.ResponseWriter, r *http.Request) {
	var ish = &entity.Issue{}
	utilities.ParseBodyTest(r, ish, w)
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]

	if ok, errors := utilities.ValidateInputs(ish); !ok {
		utilities.ValidationResponse(errors, w)
		return
	}
	v, err := models.AddIssue(ish, webpageId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(v, w)

}

func GetAllIssuesforWebpageId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	issues, err := models.GetAllIssuesforWebpageId(webpageId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	res, _ := json.Marshal(issues)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetIssues(w http.ResponseWriter, r *http.Request) {
	webpageScan, err := models.GetAllWebpages()
	if err != nil {

	}
	res, _ := json.Marshal(webpageScan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//func GetWebpageScanById(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	webpageId := vars["webpageId"]
//	webpageDetails, err := models.GetWebpageById(webpageId)
//	if err != nil {
//
//	}
//	res, _ := json.Marshal(webpageDetails)
//	w.Header().Set("Content-Type", "pkglication/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}

func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var updateWebpage = &entity.Webpage{}
	utilities.ParseBody(r, updateWebpage)
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]

	webpageDetails, err := models.GetWebpageById(webpageId)
	if err != nil {
		println("tomiiiii", err)
		return

	}

	if updateWebpage.Name != "" {
		webpageDetails.Name = updateWebpage.Name
	}
	if updateWebpage.Note != "" {
		webpageDetails.Note = updateWebpage.Note
	}
	models.UpdateWebpage(webpageDetails, webpageId)
	res, _ := json.Marshal(webpageDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteIssueByWebpageAndWebpageId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	issueId := vars["issueId"]
	wpageId, err := primitive.ObjectIDFromHex(webpageId)
	ishId, err := primitive.ObjectIDFromHex(issueId)
	if err != nil {
		utilities.ErrorResponse(422, err.Error(), w)
		return
	}
	vin, error := models.DeleteIssueByWebpageAndWebpageId(wpageId, ishId)
	if error != nil {
		utilities.ErrorResponse(500, error.Error(), w)
		return
	}
	res, _ := json.Marshal(vin)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
