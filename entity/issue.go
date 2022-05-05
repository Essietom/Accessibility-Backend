package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Issue struct {
	ID        primitive.ObjectID `json:"issueId" bson:"_id,omitempty"`
	Criteria  []Criteria         `json:"criteria" validate:"required" `
	Occurence   []Occurence          `json:"occurences" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Timestamp string             `json:"timestamp"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
	Name      string             `json:"name"`
}


