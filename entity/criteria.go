package entity

type Criteria struct {
	ID   string `json:"criteriaId" bson:"_id,omitempty"`
	Name string `json:"name" validate:"required"`
	Note string `json:"note"`
}
