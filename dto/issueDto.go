package dto

import "Accessibility-Backend/entity"

type IssueRequestBody struct {
	Criteria  []entity.Criteria         `json:"criteria" validate:"required" `
	Finding   []entity.Finding          `json:"finding" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
}


type IssueResponseBody struct {
	ID        string `json:"issueId"`
	Criteria  []Criteria         `json:"criteria"`
	Finding   []Finding          `json:"finding"`
	Impact    string             `json:"impact"`
	Timestamp string             `json:"timestamp"`
	Found     string             `json:"found"`
	Note      string             `json:"note"`
}
