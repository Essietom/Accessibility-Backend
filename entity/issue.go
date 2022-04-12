package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Issue struct {
	ID        primitive.ObjectID `json:"issueId" bson:"_id,omitempty"`
	Criteria  []Criteria         `json:"criteria" validate:"required" `
	Finding   []Finding          `json:"finding" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Timestamp string             `json:"timestamp"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
}


type IssueDto struct {
	ID        string `json:"issueId"`
	Criteria  []Criteria         
	Finding   []Finding    
	Impact    string          
	Timestamp string           
	Found     string           
	Note      string             
}

type IssueInsert struct {
	ID        string `json:"issueId"`
	Criteria  []Criteria         `json:"criteria" validate:"required" `
	Finding   []Finding          `json:"finding" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Timestamp string             `json:"timestamp"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
}

type IssueUpdate struct {
	ID        string `json:"issueId"`
	Criteria  []Criteria         `json:"criteria" validate:"required" `
	Finding   []Finding          `json:"finding" validate:"required"`
	Impact    string             `json:"impact" validate:"required"`
	Timestamp string             `json:"timestamp"`
	Found     string             `json:"found" validate:"required"`
	Note      string             `json:"note"`
}



// type UpdateUser struct {
//     FirstName string    `json:"firstName"`
//     LastName  string    `json:"lastName"`
//     ... fields that are updatable with appropriate validation tags
// }

// func (u *UpdateUser) ToModel() *model.User {
//     return &model.User{
//        FirstName: u.FirstName,
//        LastName: u.LastName,
//        ...
//     }
// }