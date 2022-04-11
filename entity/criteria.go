package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Criteria struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" validate:"required"`
	Note string             `json:"note"`
}
