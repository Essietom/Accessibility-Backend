package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Website struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" validate:"required"`
	Url  string             `json:"url" validate:"required"`
}
