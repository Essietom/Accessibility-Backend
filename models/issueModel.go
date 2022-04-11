package models

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//update website, add issue to issue object
func AddIssue(issue *entity.Issue, websiteId string) (*entity.Issue, error) {

	if issue.Criteria == nil {
		issue.Criteria = make([]entity.Criteria, 0) // this is alloc free
	}
	if issue.Finding == nil {
		issue.Finding = make([]entity.Finding, 0) // this is alloc free
	} else {
		issue.Finding[0].ID = primitive.NewObjectID()
	}

	objectId, err := primitive.ObjectIDFromHex(websiteId)

	issue.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	issue.ID = primitive.NewObjectID()

	result, err := database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": objectId},
		bson.M{
			"$push": bson.M{
				"issue": &issue,
			},
		},
	)
	if err != nil {
		return nil, errors.New("some error occurred while entering issue")
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("webpage not found")
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("issue could not be updated")
	}

	return issue, err

}

//return all the issues for a webpage
func GetAllIssues() ([]entity.Webpage, error) {
	var webpage entity.Webpage
	var webpages []entity.Webpage
	cursor, err := database.WebpageCollection.Find(database.Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(database.Ctx)
		return webpages, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&webpage)
		if err != nil {
			return webpages, err
		}
		webpages = append(webpages, webpage)
	}
	return webpages, nil
}

func GetAllIssuesforWebpageId(id string) ([]entity.Issue, error) {
	var wp entity.Webpage
	var issues []entity.Issue
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("the webpage id passed is invalid")
	}
	err = database.WebpageCollection.
		FindOne(database.Ctx, bson.D{{"_id", objectId}}).
		Decode(&wp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no Issues were found for this webpage")
		}
		return nil, err
	}
	for _, element := range wp.Issue {
		issues = append(issues, element)
	}

	return issues, nil
}
func GetIssueByIssueIdAndWebpageId(issueId string, webpageId string) (*entity.Issue, error) {
	var issue entity.Issue
	ishId, err := primitive.ObjectIDFromHex(issueId)
	wpageId, err := primitive.ObjectIDFromHex(webpageId)
	fmt.Println("issueId", ishId)
	fmt.Println("webpageId", wpageId)
	if err != nil {
		return nil, err
	}
	cursor, err := database.WebpageCollection.
		Aggregate(database.Ctx, bson.A{
			bson.M{
				"$match": bson.M{
					"_id": wpageId,
				},
			},
			bson.M{
				"$unwind": "$issue",
			},
			bson.M{
				"$match": bson.M{
					"issue._id": ishId,
				},
			},
			bson.M{
				"$replaceRoot": bson.M{
					"newRoot": "$issue",
				},
			},
		},
		)

	fmt.Println("cursor", cursor)

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&issue)
		if err != nil {
			fmt.Println("final isah", &issue)
			return &issue, err

		}
	}
	fmt.Println("no error ish", &issue)
	return &issue, nil
}

//update website websiteid, pull issue where webpage id = webpage and issue id = issueid
func DeleteIssueByWebpageAndWebpageId(webpageId primitive.ObjectID, issueId primitive.ObjectID) (string, error) {
	result, err := database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": webpageId},
		bson.M{
			"$pull": bson.M{
				"issue": bson.M{
					"_id": issueId,
				},
			},
		},
	)
	if err != nil {
		return "nil", errors.New("some error occurred")
	}
	if result.MatchedCount == 0 {
		return "nil", errors.New("no webpage found with the id provided")
	}
	if result.ModifiedCount == 0 {
		return "nil", errors.New("no issue found for provided id,issue was not successfully updated")
	}

	return "Issue was successfully deleted", nil
}

//db.users.update({_id:123}, {$set:{
//"friends.$[updateFriend].emails.$[updateEmail].email : "lucy.is.gucy@zmail.com"
//}}, {
//"arrayFilters": [
//{"updateFriend.name" : "lucy"},
//{"updateEmail.email" : "lucyGucy@zmail.com"}
//]
//})
//db.users.update ({_id: '123'}, { '$set': {"friends.0.emails.0.email" : '2222'} });
//updates an issue under a website and the findings under that issue
func UpdateIssue(v *entity.Webpage, id string) (*entity.Webpage, error) {

	result, err := database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": id},
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
	v, err = GetWebpageById(id)
	if err != nil {
		return nil, err
	}
	return v, err
}
