package dto

import (
	"Accessibility-Backend/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Impact int
type Found int

const (
	Automatically Found = iota
	Manually
)
// String - Creating common behavior - give the type a String function
func (f Found) String() string {
	return [...]string{"Automatically", "Manually"}[f]
}
const (
	Minor Impact = iota
	Moderate
	Serious
	Critical
)
func (i Impact) String() string {
	return [...]string{"Minor", "Moderate", "Serious", "Critical"}[i]
}

type IssueRequestBody struct {
	Criteria  []CriteriaRequestBody         `json:"criteria" validate:"required" `
	Occurence   []OccurenceRequestBody          `json:"occurences" validate:"required"`
	Impact    Impact             `json:"impact" validate:"required"`
	Found     Found             `json:"found" validate:"required"`
	Note      string             `json:"note"`
	Name      string             `json:"name"`
}


type IssueUpdateBody struct {
	Occurence   []OccurenceRequestBody         `json:"occurences"`
	Impact    string             `json:"impact"`
	Note      string             `json:"note"`
}


func (data IssueRequestBody) ToIssueEntities() *entity.Issue {
	return &entity.Issue{
		ID: primitive.NewObjectID(),
		Impact:        data.Impact.String(),
		Found:       data.Found.String(),
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