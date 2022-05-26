package repository

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func FindOccurenceById(issueId primitive.ObjectID, webpageId primitive.ObjectID, occurenceId primitive.ObjectID) (*mongo.Cursor, error) {
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
				"$unwind": "$issue.occurence",
			},
			bson.M{
				"$match": bson.M{
					"issue.occurence._id": occurenceId,
				},
			},
			bson.M{
				"$replaceRoot": bson.M{
					"newRoot": "$issue.occurence",
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

func GetOccurenceCount(issueId primitive.ObjectID, webpageId primitive.ObjectID) (*mongo.Cursor, error) {
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
				"$project": bson.M{
					"count": bson.M{
						"$size": "$issue.occurence",
					},
				},
			},
		},
		)
}
			


func DeleteOccurence(webpageId primitive.ObjectID, issueId primitive.ObjectID, occurenceId primitive.ObjectID) (*mongo.UpdateResult, error) {
	return database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": webpageId, "issue._id":issueId},
		bson.M{
			"$pull": bson.M{
				"issue.$.occurence": bson.M{
					"_id": occurenceId,
				},
			},
		},
	)

	
}

func UpdateIssue(issueBody *entity.Issue, webpageId primitive.ObjectID, issueId primitive.ObjectID) (*mongo.SingleResult) {

	return database.WebpageCollection.FindOneAndUpdate(database.Ctx, bson.M{"_id": webpageId},
			bson.M{
				"$set": bson.M{
					"issue.$[elem]": &issueBody,
				},
			},
			options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
				Filters: []interface{}{bson.M{"elem._id": issueId}},
			}).SetReturnDocument(1),
			
		)

	
}

func UpdateOccurence(occurenceBody *entity.Occurence, webpageId primitive.ObjectID, issueId primitive.ObjectID, occurenceId primitive.ObjectID) (*mongo.SingleResult) {

	return database.WebpageCollection.FindOneAndUpdate(database.Ctx, bson.M{"_id": webpageId},
			bson.M{
				"$set": bson.M{
					"issue.$[elem].occurence.$[elem2]": &occurenceBody,
				},
			},
			options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
				Filters: []interface{}{bson.M{"elem._id": issueId},
				bson.M{
					"elem2._id": occurenceId,
				},},
			}).SetReturnDocument(1),
			
		)

	
}