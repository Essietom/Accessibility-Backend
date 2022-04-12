package controllers

import (
	"Accessibility-Backend/entity"
	"Accessibility-Backend/models"
	"Accessibility-Backend/utilities"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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




func UpdateIssueByIssueIdAndWebpageId(w http.ResponseWriter, r *http.Request) {
	var updateIssue = &entity.Issue{}
	utilities.ParseBody(r, updateIssue)
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	issueId := vars["issueId"]

	wpageId, err := primitive.ObjectIDFromHex(webpageId)
	ishId, err := primitive.ObjectIDFromHex(issueId)
	if err != nil {
		utilities.ErrorResponse(422, err.Error(), w)
		return
	}

	issueDetails, err := models.GetIssueByIssueIdAndWebpageId(ishId, wpageId)
	if err != nil {
		utilities.ErrorResponse(404, "no issue with the provided id", w)
		return

	}

	if updateIssue.Note != "" {
		issueDetails.Note = updateIssue.Note
	}
	if updateIssue.Impact != "" {
		issueDetails.Impact = updateIssue.Impact
	}
	if updateIssue.Finding != nil && updateIssue.Finding[0].Description != "" {
		issueDetails.Finding[0].Description = updateIssue.Finding[0].Description
	}
	res, err := models.UpdateIssueByIssueIdAndWebpageId(issueDetails, wpageId, ishId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(res, w)

}
func DeleteIssueByIssueIdAndWebpageId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	issueId := vars["issueId"]
	wpageId, err := primitive.ObjectIDFromHex(webpageId)
	ishId, err := primitive.ObjectIDFromHex(issueId)
	if err != nil {
		utilities.ErrorResponse(422, err.Error(), w)
		return
	}
	vin, error := models.DeleteIssueByIssueIdAndWebpageId(wpageId, ishId)
	if error != nil {
		utilities.ErrorResponse(500, error.Error(), w)
		return
	}
	res, _ := json.Marshal(vin)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
