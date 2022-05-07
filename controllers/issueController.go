package controllers

import (
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"net/http"
)

var Issue entity.Issue


func AddIssue(w http.ResponseWriter, r *http.Request) {

	var ish = &dto.IssueRequestBody{}
	utilities.ParseBodyTest(r, ish, w)
	webpageId := r.URL.Query().Get("webpageId")


	if ok, errors := utilities.ValidateInputs(ish); !ok {
		utilities.ValidationResponse(errors, w, r)
		return
	}
	v, err := model.AddIssue(*ish, webpageId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(v, w, r)

}

func GetAllIssuesforWebpageId(w http.ResponseWriter, r *http.Request) {

	webpageId := r.URL.Query().Get("webpageId")

	 issues, err := model.GetIssuesByWebpageId(webpageId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(issues, w, r)
}




func UpdateIssueByIssueIdAndWebpageId(w http.ResponseWriter, r *http.Request) {

	var updateIssue = &dto.IssueUpdateBody{}
	utilities.ParseBody(r, updateIssue)

	issueId := r.URL.Query().Get("issueId")
	webpageId := r.URL.Query().Get("webpageId")
	issueDetails, err := model.GetIssueByWebpageIdAndIssueId(issueId, webpageId)
	if err != nil {
		utilities.ErrorResponse(404, "no issue with the provided id", w, r)
		return

	}


	if updateIssue.Note != "" {
		issueDetails.Note = updateIssue.Note
	}
	if updateIssue.Impact != "" {
		issueDetails.Impact = updateIssue.Impact
	}
	if updateIssue.Occurence != nil && updateIssue.Occurence[0].Description != "" {
		issueDetails.Occurence[0].Description = updateIssue.Occurence[0].Description
	}

	 res, err := model.UpdateIssueByIssueIdAndWebpageId(issueDetails, webpageId, issueId)

	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(res, w, r)

}
func DeleteIssueByIssueIdAndWebpageId(w http.ResponseWriter, r *http.Request) {

	issueId := r.URL.Query().Get("issueId")
	webpageId := r.URL.Query().Get("webpageId")
	
	error := model.DeleteIssue(webpageId, issueId)

	if error != nil {
		utilities.ErrorResponse(500, error.Error(), w, r)
		return
	}
	res := "successfully deleted"
	utilities.SuccessRespond(res, w, r)

}

func DeleteOccurenceIdAndIssueIdAndWebpageId(w http.ResponseWriter, r *http.Request) {

	issueId := r.URL.Query().Get("issueId")
	webpageId := r.URL.Query().Get("webpageId")
	occurenceId := r.URL.Query().Get("occurenceId")

	error := model.DeleteOccurence(webpageId, issueId, occurenceId)

	if error != nil {
		utilities.ErrorResponse(500, error.Error(), w, r)
		return
	}
	res := "successfully deleted"
	utilities.SuccessRespond(res, w, r)

}
