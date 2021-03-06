package dto

import (
	"Accessibility-Backend/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Impact int
type Found int

// const (
// 	automatic Found = iota
// 	guided
// 	needsReview
// 	manual
// )
// String - Creating common behavior - give the type a String function
func (f Found) String() string {
	return [...]string{"automatic", "guided", "needsReview", "manual"}[f]
}
// const (
// 	minor Impact = iota
// 	moderate    
// 	serious
// 	critical
// )
func (i Impact) String() string {
	return [...]string{"minor", "moderate", "serious", "critical"}[i]
}

type IssueRequestBody struct {
	Criteria  []CriteriaRequestBody         `json:"criteria" validate:"required" `
	Occurence   []OccurenceRequestBody          `json:"occurences" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
	Name      string             `json:"name"`
}


type IssueUpdateBody struct {
	Occurence   []OccurenceRequestBody         `json:"occurences"`
	Impact    string             `json:"impact"`
	Note      string             `json:"note"`
	Name      string             `json:"name"`
	Criteria  []entity.Criteria         `json:"criteria"`

}


func (data IssueRequestBody) ToIssueEntities() *entity.Issue {
	return &entity.Issue{
		ID: primitive.NewObjectID(),
		Impact:        data.Impact,
		Found:       data.Found,
		Note: data.Note,
		Name: data.Name,
		Criteria:    *GetIssueCriteria(data.Criteria),
		Occurence:    *GetOccurenceEntities(data.Occurence),
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