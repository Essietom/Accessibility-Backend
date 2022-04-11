package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddIssueRequest struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Criteria  []Criteria         `json:"criteria" validate:"required" `
	Finding   []Finding          `json:"finding" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Timestamp string             `json:"timestamp"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
}

type Issue struct {
	//ID string `json:"id" bson:"_id,omitempty"`
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Criteria  []Criteria         `json:"criteria" validate:"required" `
	Finding   []Finding          `json:"finding" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Timestamp string             `json:"timestamp"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
}
