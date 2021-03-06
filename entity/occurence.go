package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Occurence struct {
	//ID string `json:"id" bson:"_id,omitempty"`

	ID          primitive.ObjectID `json:"occurenceId" bson:"_id,omitempty"`
	Description string             `json:"description" validate:"required"`
	Location    string             `json:"location"`
	Source      string             `json:"source"`
	Fix         string             `json:"fix"`
	Note        string             `json:"note"`
	NeedsReview        bool             `json:"needsReview"`

}
