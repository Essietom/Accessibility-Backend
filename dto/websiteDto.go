package dto

import (
	"Accessibility-Backend/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebsiteRequestBody struct {
	Name string             `json:"name" validate:"required"`
	Url  string             `json:"url" validate:"required"`
}

func (data WebsiteRequestBody) ToWebsiteEntities() *entity.Website {
	return &entity.Website{
		ID: primitive.NewObjectID(),
		Name:       data.Name,
		Url:       data.Url,
		
	}
}