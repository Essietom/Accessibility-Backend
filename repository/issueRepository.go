package repository

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func FindByWebpageId(webpageId primitive.ObjectID) ([]entity.Issue, error) {
	var wp entity.Webpage
	err := database.WebpageCollection.
		FindOne(database.Ctx, bson.D{{"_id", webpageId}}).
		Decode(&wp)
	return wp.Issue, err
}

func FindByIssueIdAndWebpageId(issueId primitive.ObjectID, webpageId primitive.ObjectID) (*mongo.Cursor, error) {
	return database.WebpageCollection.
		Aggregate(database.Ctx, bson.A{
			bson.M{
				"$match": bson.M{
					"_id": webpageId,
				},
			},
			bson.M{
				"$unwind": "$issue",
			},
			bson.M{
				"$match": bson.M{
					"issue._id": issueId,
				},
			},
			bson.M{
				"$replaceRoot": bson.M{
					"newRoot": "$issue",
				},
			},
		},
		)
}

func SaveIssue(issue *entity.Issue, websiteId primitive.ObjectID) (*mongo.UpdateResult, error) {
	return database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": websiteId},
		bson.M{
			"$push": bson.M{
				"issue": &issue,
			},
		},
	)
}

func DeleteIssue(webpageId primitive.ObjectID, issueId primitive.ObjectID) (*mongo.UpdateResult, error) {

	return database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": webpageId},
		bson.M{
			"$pull": bson.M{
				"issue": bson.M{
					"_id": issueId,
				},
			},
		},
	)

	
}