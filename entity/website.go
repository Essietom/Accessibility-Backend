package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Website struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name"`
	Url  string             `json:"url"`
}
