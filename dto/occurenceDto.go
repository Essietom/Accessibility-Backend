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
}

func (data OccurenceRequestBody) ToOccurenceEntities() *entity.Occurence {
	return &entity.Occurence{
		ID: primitive.NewObjectID(),
		Description:       data.Description,
		Location:       data.Location,
		Source:       data.Source,
		Fix:       data.Fix,
		Note: data.Note,
	}
}

func  GetOccurenceEntities(occurences []OccurenceRequestBody) *[]entity.Occurence{
	var occurenceEntities []entity.Occurence
	for _, occurrence := range occurences {	
		occurenceEntities = append(occurenceEntities, *occurrence.ToOccurenceEntities())
	}
	return &occurenceEntities
}