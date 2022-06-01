package controllers

import (
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"log"
	"net/http"
)

var Issue entity.Issue
var (
	T = true
	F = false
)

func AddIssue(w http.ResponseWriter, r *http.Request) {

	var ish = &dto.IssueRequestBody{}
	err := utilities.ParseBodyTest(r, ish, w)
	 if err != nil {
		utilities.ErrorResponse(http.StatusBadRequest, err.Error(), w,r)
		return
	}
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
	occurenceId := r.URL.Query().Get("occurenceId")

	

	if(updateIssue.Occurence != nil){
		var updateOccu = &dto.OccurenceUpdateBody{}

		if updateIssue.Occurence[0].Description != "" {
			updateOccu.Description = updateIssue.Occurence[0].Description
		}
		if updateIssue.Occurence[0].Note != "" {
			updateOccu.Note = updateIssue.Occurence[0].Note
		}
		if updateIssue.Occurence[0].Fix != "" {
			updateOccu.Fix = updateIssue.Occurence[0].Fix
		}
		
		if updateIssue.Occurence[0].NeedsReview != nil {
			if !(*updateIssue.Occurence[0].NeedsReview) {
				updateOccu.NeedsReview = &F;
			} else {
				updateOccu.NeedsReview = &T;
			}
			
		}  
		errCode, errMessage := UpdateOccurence2(issueId, webpageId, occurenceId, *updateOccu)
		if errCode!=200{
			utilities.ErrorResponse(errCode, errMessage, w, r)
			return
		}
	}

	issueDetails, err := model.GetIssueByWebpageIdAndIssueId(issueId, webpageId)
	if err != nil {
		utilities.ErrorResponse(404, "no issue with the provided id", w, r)
		return

	}

	if updateIssue.Name != "" {
		issueDetails.Name = updateIssue.Name
	}
	if updateIssue.Criteria != nil {
		issueDetails.Criteria = updateIssue.Criteria
	}
	if updateIssue.Note != "" {
		issueDetails.Note = updateIssue.Note
	}
	if updateIssue.Impact != "" {
		issueDetails.Impact = updateIssue.Impact
	}

	
	res, err := model.UpdateIssueByIssueIdAndWebpageId(issueDetails, webpageId, issueId)


	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(res, w, r)

}

func UpdateOccurence2(issueId string, webpageId string, occurenceId string, updateOccurrence dto.OccurenceUpdateBody) (int, string){

	occurenceDetails, err := model.GetOccurenceById(issueId, webpageId, occurenceId)
	if err != nil {
		return 404, "no occurence with the provided id"
	}


	if updateOccurrence.Note != "" {
		occurenceDetails.Note = updateOccurrence.Note
	}
	if updateOccurrence.Description != "" {
		occurenceDetails.Description = updateOccurrence.Description
	}
	if updateOccurrence.Fix != "" {
		occurenceDetails.Fix = updateOccurrence.Fix
	}
	if updateOccurrence.NeedsReview != nil {
		if !(*updateOccurrence.NeedsReview) {
			occurenceDetails.NeedsReview = F;
		} else {
			occurenceDetails.NeedsReview = T;
		}
        
    }  

	 res, err := model.UpdateOccurence(occurenceDetails, webpageId, issueId, occurenceId)
	if err != nil {
		return 500, err.Error()
	}
	return 200, res.ID.Hex()

}


func UpdateOccurence(w http.ResponseWriter, r *http.Request) {

	var updateOccurrence = &dto.OccurenceUpdateBody{}
	utilities.ParseBody(r, updateOccurrence)

	issueId := r.URL.Query().Get("issueId")
	webpageId := r.URL.Query().Get("webpageId")
	occurenceId := r.URL.Query().Get("occurenceId")

	occurenceDetails, err := model.GetOccurenceById(issueId, webpageId, occurenceId)
	log.Println("found this:", occurenceDetails)
	if err != nil {
		utilities.ErrorResponse(404, "no occurence with the provided id", w, r)
		return

	}


	if updateOccurrence.Note != "" {
		occurenceDetails.Note = updateOccurrence.Note
	}
	if updateOccurrence.Description != "" {
		occurenceDetails.Description = updateOccurrence.Description
	}
	if updateOccurrence.Fix != "" {
		occurenceDetails.Fix = updateOccurrence.Fix
	}

	 res, err := model.UpdateOccurence(occurenceDetails, webpageId, issueId, occurenceId)

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

	error := model.DeleteOccurenceOrIssue(webpageId, issueId, occurenceId)

	if error != nil {
		utilities.ErrorResponse(500, error.Error(), w, r)
		return
	}
	res := "successfully deleted"
	utilities.SuccessRespond(res, w, r)

}
