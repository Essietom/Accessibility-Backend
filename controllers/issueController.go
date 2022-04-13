package controllers

import (
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

var Issue entity.Issue


func AddIssue(w http.ResponseWriter, r *http.Request) {
	var ish = &dto.IssueRequestBody{}
	utilities.ParseBodyTest(r, ish, w)
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]

	if ok, errors := utilities.ValidateInputs(ish); !ok {
		utilities.ValidationResponse(errors, w)
		return
	}
	v, err := model.AddIssue(*ish, webpageId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(v, w)

}

func GetAllIssuesforWebpageId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	 issues, err := model.GetIssuesByWebpageId(webpageId)
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


	issueDetails, err := model.GetIssueByWebpageIdAndIssueId(issueId, webpageId)
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
	// res, err := model.UpdateIssueByIssueIdAndWebpageId(issueDetails, wpageId, ishId)
	res, err := "nil",nil

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
	
	error := model.DeleteIssue(webpageId, issueId)

	if error != nil {
		utilities.ErrorResponse(500, error.Error(), w)
		return
	}
	res, _ := json.Marshal("successfully deleted")
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
