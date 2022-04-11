package entity

type Criteria struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" validate:"required"`
	Note string `json:"note"`
}
