package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Finding struct {
	//ID string `json:"id" bson:"_id,omitempty"`

	ID          primitive.ObjectID `json:"findingId" bson:"_id,omitempty"`
	Description string             `json:"description" validate:"required"`
	Location    string             `json:"location"`
	Source      string             `json:"source"`
	Fix         string             `json:"fix"`
	Note        string             `json:"note"`
}
