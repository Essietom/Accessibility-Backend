package dto

import (
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindingRequestBody struct {
	Description string             `json:"description" validate:"required"`
	Location    string             `json:"location"`
	Source      string             `json:"source"`
	Fix         string             `json:"fix"`
	Note        string             `json:"note"`
}

func (data FindingRequestBody) ToFindingEntities() *entity.Finding {
	return &entity.Finding{
		ID: primitive.NewObjectID(),
		Description:       data.Description,
		Location:       data.Location,
		Source:       data.Source,
		Fix:       data.Fix,
		Note: data.Note,
	}
}

func  GetFindingEntities(findings []FindingRequestBody) *[]entity.Finding{
	var findingEntities []entity.Finding
	for _, findi := range findings {	
		findingEntities = append(findingEntities, *findi.ToFindingEntities())
	}
	return &findingEntities
}