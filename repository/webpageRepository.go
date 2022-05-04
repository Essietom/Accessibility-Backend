package repository

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetWebpageByField(queryField string, sortField string, orderby int, pageSize int64, pageNum int64) (*mongo.Cursor, error){
	findOptions := options.Find()
	skips := pageSize * (pageNum - 1)
	findOptions.SetSort(bson.D{{sortField, orderby}})
	findOptions.SetSkip(skips)
	findOptions.SetLimit(pageSize)
	return database.WebpageCollection.
		Find(database.Ctx,
			bson.M{
				"$or": bson.A{
					bson.M{
						"name": &queryField,
					},
					bson.M{
						"website.name": &queryField,
					},
					bson.M{
						"url": &queryField,
					},
					bson.M{
						"website.url": &queryField,
					},
				},
			},
		findOptions)
}