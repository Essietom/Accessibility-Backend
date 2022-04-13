package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Webpage struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" `
	Url      string             `json:"url" `
	ScanTime string             `json:"scanTime"`
	Note     string             `json:"note"`
	Issue    []Issue            `json:"issue" `
	Website  Website            `json:"website"`
}
