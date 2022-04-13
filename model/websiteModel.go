package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateWebsite(v *entity.Website) (*entity.Website, error) {
	result, err := repository.SaveWebsite(v)
	if err != nil {
		fmt.Println("unable to insert record", err)
		return nil, err
	}
	v.ID = result.InsertedID.(primitive.ObjectID)
	return v, err
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
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
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

func UpdateWebsite(v *entity.Website, id string) (*entity.Website, error) {

	objectId, _ := primitive.ObjectIDFromHex(id)
	result, err := repository.UpdateWebsite(v, objectId)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, err
	}
	v, err = GetWebsiteById(id)
	if err != nil {
		return nil, err
	}
	return v, err
}
