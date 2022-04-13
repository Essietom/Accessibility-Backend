package repository

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func FindWebpageById(webpageId primitive.ObjectID) (entity.Webpage, error) {
	var wp entity.Webpage
	err := database.WebpageCollection.
		FindOne(database.Ctx, bson.D{{"_id", webpageId}}).
		Decode(&wp)
	return wp, err
}

func FindWebpages() (*mongo.Cursor, error) {
	return database.WebpageCollection.Find(database.Ctx, bson.D{})
}

func SaveWebpage(wp *entity.Webpage) (*mongo.InsertOneResult, error) {
	return database.WebpageCollection.InsertOne(database.Ctx, wp)
}

func DeleteWebpage(id primitive.ObjectID) error {

	_, err := database.WebpageCollection.DeleteOne(database.Ctx, bson.D{{"_id", id}})
	return err

	
}

// func UpdateWebpage(v *entity.Webpage, id primitive.ObjectID) (*mongo.UpdateResult, error) {

// 	return database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": id},
// 	bson.M{
// 		"$set": &v,
// 	},
// )

	
// }

func GetWebpageByField(field string) (*mongo.Cursor, error){
	return database.WebpageCollection.
		Find(database.Ctx,
			bson.M{
				"$or": bson.A{
					bson.M{
						"name": &field,
					},
					bson.M{
						"website.name": &field,
					},
					bson.M{
						"url": &field,
					},
					bson.M{
						"website.url": &field,
					},
				},
			})
}