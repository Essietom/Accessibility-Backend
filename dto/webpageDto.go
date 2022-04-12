package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Webpage struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" validate:"required"`
	Url      string             `json:"url" validate:"required"`
	ScanTime string             `json:"scanTime" validate:"required"`
	Note     string             `json:"note"`
	// Issue    []Issue            `json:"issue" validate:"required"`
	Website  Website            `json:"website" validate:"required"`
}
