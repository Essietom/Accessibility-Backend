package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


//update website, add issue to issue object
func AddIssue(data dto.IssueRequestBody, websiteId string) (*entity.Issue, error) {

	if data.Criteria == nil {
		data.Criteria = make([]entity.Criteria, 0) // this is alloc free
	}
	if data.Finding == nil {
		data.Finding = make([]entity.Finding, 0) // this is alloc free
	} else {
		data.Finding[0].ID = primitive.NewObjectID()
	}

	isss := entity.Issue{
		Impact:        data.Impact,
		Found:       data.Found,
		Note: data.Note,
		Criteria:    data.Criteria,
	}

	objectId, _ := primitive.ObjectIDFromHex(websiteId)

	isss.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	isss.ID = primitive.NewObjectID()

	result, err := repository.Save(&isss, objectId)
	if err != nil {
		return nil, errors.New("some error occurred while entering issue")
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("webpage not found")
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("issue could not be updated")
	}

	return &isss, err

}



func GetIssuesByWebpageId(id string) ([]entity.Issue, error) {
	var issues []entity.Issue
	objectId, _ := primitive.ObjectIDFromHex(id)
	
	ishs, err := repository.FindByWebpageId(objectId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no Issues were found for this webpage")
		}
		return nil, err
	}
	for _, element := range ishs {
		issues = append(issues, element)
	}

	return issues, nil
}

func GetIssueByWebpageIdAndIssueId(issueId string, webpageId string) (*entity.Issue, error) {
	var issue entity.Issue
	ishId, _ := primitive.ObjectIDFromHex(issueId)	
	wpId, _ := primitive.ObjectIDFromHex(webpageId)


	cursor, err := repository.FindByIssueIdAndWebpageId(ishId, wpId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no Issues were found for this webpage")
		}
		return nil, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&issue)
		if err != nil {
			return &issue, err

		}
	}
	return &issue, nil
}


//update website websiteid, pull issue where webpage id = webpage and issue id = issueid
func DeleteIssue(webpageId string, issueId string) error {

	ishId, _ := primitive.ObjectIDFromHex(issueId)	
	wpId, _ := primitive.ObjectIDFromHex(webpageId)

	result, err := repository.Delete(wpId, ishId)
	if err != nil {
		return errors.New("some error occurred")
	}
	if result.MatchedCount == 0 {
		return  errors.New("no webpage found with the id provided")
	}
	if result.ModifiedCount == 0 {
		return errors.New("no issue found for provided id,issue was not successfully updated")
	}

	return  nil
}

//updates an issue under a website and the findings under that issue
//todo not working
// func UpdateIssueByIssueIdAndWebpageId(v *entity.Issue, webpageId primitive.ObjectID, issueId primitive.ObjectID) (*entity.Issue, error) {

// 	result, err := database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": webpageId},
// 		bson.M{
// 			"$set": bson.M{
// 				"issue.$": &v,
// 				"arrayFilters": bson.A{
// 					bson.M{
// 						"issue._id": issueId,
// 					},
// 				},
// 			},
			
// 		},
// 	)


// 	if err != nil {
// 		fmt.Println("what", err)
// 		return nil, errors.New("some error occurred while updating issue")
// 	}
// 	if result.MatchedCount == 0 {
// 		return nil, errors.New("webpage not found")
// 	}
// 	if result.ModifiedCount == 0 {
// 		return nil, errors.New("issue could not be updated")
// 	}

// 	return v, err
// }
