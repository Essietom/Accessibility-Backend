package repository

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindWebsiteById(website primitive.ObjectID) (entity.Website, error) {
	var ws entity.Website
	err := database.WebsiteCollection.
	FindOne(database.Ctx, bson.D{{"_id", website}}).
	Decode(&ws)
	return ws, err
}

func FindWebsites() (*mongo.Cursor, error) {
	return database.WebsiteCollection.Find(database.Ctx, bson.D{})
}

func SaveWebsite(wp *entity.Website) (*mongo.InsertOneResult, error) {
	return database.WebsiteCollection.InsertOne(database.Ctx, wp)
}



func UpdateWebsite(v *entity.Website, id primitive.ObjectID) (*mongo.UpdateResult, error) {

	return database.WebsiteCollection.UpdateOne(database.Ctx, bson.M{"_id": id},
	bson.M{
		"$set": &v,
	},
)

	
}

func GetWebsiteByField(field string) (entity.Website, error){
	var ws entity.Website
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
	Decode(&ws)
	return ws, err
}