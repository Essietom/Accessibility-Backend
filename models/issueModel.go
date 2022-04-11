package models

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//update website, add issue to issue object
func AddIssue(issue *entity.Issue, websiteId string) (string, error) {
	fmt.Println("the ish", issue)

	if issue.Criteria == nil {
		issue.Criteria = make([]entity.Criteria, 0) // this is alloc free
	}
	if issue.Finding == nil {
		issue.Finding = make([]entity.Finding, 0) // this is alloc free
	} else {
		issue.Finding[0].ID = primitive.NewObjectID()
	}

	objectId, err := primitive.ObjectIDFromHex(websiteId)
	issue.ID = primitive.NewObjectID()
	result, err := database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": objectId},
		bson.M{
			"$push": bson.M{
				"issue": &issue,
			},
		},
	)
	if err != nil {
		return "error occured", err
	}
	if result.MatchedCount == 0 {
		return "no match found", err
	}
	if result.ModifiedCount == 0 {
		return "not updated", err
	}

	return "successful", err

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
		return nil, err
	}
	err = database.WebpageCollection.
		FindOne(database.Ctx, bson.D{{"_id", objectId}}).
		Decode(&wp)

	if err != nil {
		return nil, err
	}
	for _, element := range wp.Issue {
		issues = append(issues, element)
	}

	return issues, nil
}

//update website websiteid, pull issue where issue id = issueid
func DeleteIssue(id primitive.ObjectID) error {
	_, err := database.WebpageCollection.DeleteOne(database.Ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return err
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
