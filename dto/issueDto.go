package dto

import (
	"Accessibility-Backend/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IssueRequestBody struct {
	Criteria  []CriteriaRequestBody         `json:"criteria" validate:"required" `
	Finding   []FindingRequestBody          `json:"finding" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
}


// type IssueResponseBody struct {
// 	ID        string `json:"issueId"`
// 	Criteria  []Criteria         `json:"criteria"`
// 	Finding   []Finding          `json:"finding"`
// 	Impact    string             `json:"impact"`
// 	Timestamp string             `json:"timestamp"`
// 	Found     string             `json:"found"`
// 	Note      string             `json:"note"`
// }


func (data IssueRequestBody) ToIssueEntities() *entity.Issue {
	return &entity.Issue{
		ID: primitive.NewObjectID(),
		Impact:        data.Impact,
		Found:       data.Found,
		Note: data.Note,
		Criteria:    *GetIssueCriteria(data.Criteria),
		Finding:    *GetFindingEntities(data.Finding),
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func  GetIssueEntities(issues []IssueRequestBody) *[]entity.Issue{
	var issueEntities []entity.Issue
	for _, ish := range issues {	
		issueEntities = append(issueEntities, *ish.ToIssueEntities())
	}
	return &issueEntities
}