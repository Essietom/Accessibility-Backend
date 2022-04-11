package models

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateWebsite(v *entity.Website) (*entity.Website, error) {
	result, err := database.WebsiteCollection.InsertOne(database.Ctx, v)
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
	cursor, err := database.WebsiteCollection.Find(database.Ctx, bson.D{})
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
	var v entity.Website
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = database.WebsiteCollection.
		FindOne(database.Ctx, bson.D{{"_id", objectId}}).
		Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func GetWebsiteByField(field string) (*entity.Website, error) {
	var v entity.Website

	err := database.WebsiteCollection.
		FindOne(database.Ctx,
			bson.M{
				"$or": bson.A{
					bson.M{
						"name": &field,
					},
					bson.M{
						"url": &field,
					},
				},
			}).
		Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func UpdateWebsite(v *entity.Website, id string) (*entity.Website, error) {

	result, err := database.WebsiteCollection.UpdateOne(database.Ctx, bson.M{"_id": id},
		bson.M{
			"$set": &v,
		},
	)
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
