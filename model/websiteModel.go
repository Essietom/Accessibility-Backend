package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateWebsite(websiteRequest *dto.WebsiteRequestBody) (*entity.Website, error) {

	websiteEntity := websiteRequest.ToWebsiteEntities()
	
	_, err := repository.SaveWebsite(websiteEntity)
	if err != nil {
		fmt.Println("unable to insert record", err)
		return nil, err
	}
	return websiteEntity, err
}

func GetAllWebsites() ([]entity.Website, error) {
	var website entity.Website
	var websites []entity.Website
	cursor, err := repository.FindWebsites()
	if err != nil {
		defer cursor.Close(database.Ctx)
		return websites, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&website)
		if err != nil {
			return websites, err
		}
		websites = append(websites, website)
	}
	return websites, nil
}

func GetWebsiteById(id string) (*entity.Website, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	ws, err := repository.FindWebsiteById(objectId)
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

func GetWebsiteByField(field string) (*entity.Website, error) {

	ws, err := repository.GetWebsiteByField(field)
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

func UpdateWebsite(website *entity.Website, id string) (*entity.Website, error) {

	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := repository.UpdateWebsite(website, objectId)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, err
	}
	website, err = GetWebsiteById(id)
	if err != nil {
		return nil, err
	}
	return website, err
}
