package dto

import (
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OccurenceRequestBody struct {
	Description string             `json:"description" validate:"required"`
	Location    string             `json:"location"`
	Source      string             `json:"source"`
	Fix         string             `json:"fix"`
	Note        string             `json:"note"`
	NeedsReview        *bool             `json:"needsReview"`

}

type OccurenceUpdateBody struct {
	Description string             `json:"description"`
	Location    string             `json:"location"`
	Source      string             `json:"source"`
	Fix         string             `json:"fix"`
	Note        string             `json:"note"`
	NeedsReview        *bool             `json:"needsReview"`

}

func (data OccurenceRequestBody) ToOccurenceEntities() *entity.Occurence {
	return &entity.Occurence{
		ID: primitive.NewObjectID(),
		Description:       data.Description,
		Location:       data.Location,
		Source:       data.Source,
		Fix:       data.Fix,
		Note: data.Note,
		NeedsReview: *data.NeedsReview,

	}
}

func  GetOccurenceEntities(occurences []OccurenceRequestBody) *[]entity.Occurence{
	var occurenceEntities []entity.Occurence
	for _, occurrence := range occurences {	
		occurrence.fill_defaults()
		occurenceEntities = append(occurenceEntities, *occurrence.ToOccurenceEntities())
	}
	return &occurenceEntities
}

var (
	T = true
	F = false
)
// constructor function
func(occurenceCreate *OccurenceRequestBody) fill_defaults(){
  
    // setting default values
    // if no values present
    if occurenceCreate.NeedsReview == nil {
        occurenceCreate.NeedsReview = &F
    }
      
    
}