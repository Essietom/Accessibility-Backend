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

func DeleteWebpage(id primitive.ObjectID) (*mongo.DeleteResult, error) {

	return database.WebpageCollection.DeleteOne(database.Ctx, bson.D{{"_id", id}})

	
}

// func UpdateWebpage(v *entity.Webpage, id primitive.ObjectID) (*mongo.UpdateResult, error) {

// 	return database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": id},
// 	bson.M{
// 		"$set": &v,
// 	},
// )

	
// }

func GetWebpageByField(queryField string, sortField string, orderby int, pageSize int64, pageNum int64) (*mongo.Cursor, error, int){
	filter := bson.M{}
	findOptions := options.Find()
	skips := pageSize * (pageNum - 1)
	if sortField != ""{
		findOptions.SetSort(bson.D{{sortField, orderby}})
	}
	findOptions.SetSkip(skips)
	findOptions.SetLimit(pageSize)
	total, _ := database.WebpageCollection.CountDocuments(database.Ctx, filter)

	if(queryField!=""){
		filter = bson.M{
			"$or": bson.A{
				bson.M{
					"name": bson.M{
						"$regex": primitive.Regex{
							Pattern : queryField,
							Options: "i",
						},
					},
				},
				bson.M{
					"website.name": bson.M{
						"$regex": primitive.Regex{
							Pattern : queryField,
							Options: "i",
						},
					},
				},
				bson.M{
					"url": bson.M{
						"$regex": primitive.Regex{
							Pattern : queryField,
							Options: "i",
						},
					},
				},
				bson.M{
					"website.url": bson.M{
						"$regex": primitive.Regex{
							Pattern : queryField,
							Options: "i",
						},
					},
				},
			},
		}
	}

	cur, err := database.WebpageCollection.
		Find(database.Ctx,filter,findOptions)
	return cur, err, int(total)
}