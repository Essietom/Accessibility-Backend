package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//update website, add issue to issue object
func AddIssue(data dto.IssueRequestBody, websiteId string) (*entity.Issue, error) {

	if data.Criteria == nil {
		data.Criteria = make([]dto.CriteriaRequestBody, 0) // this is alloc free
	}
	if data.Finding == nil {
		data.Finding = make([]dto.FindingRequestBody, 0) // this is alloc free
	} 
	
	isss := data.ToIssueEntities()

	objectId, _ := primitive.ObjectIDFromHex(websiteId)

	result, err := repository.SaveIssue(isss, objectId)
	if err != nil {
		return nil, errors.New("some error occurred while entering issue")
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("webpage not found")
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("issue could not be updated")
	}

	return isss, err

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

	result, err := repository.DeleteIssue(wpId, ishId)
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
func UpdateIssueByIssueIdAndWebpageId(issueUpdateBody *entity.Issue, webpageId string, issueId string) (*entity.Issue, error) {

	var issue entity.Issue
	ishId, _ := primitive.ObjectIDFromHex(issueId)	
	wpId, _ := primitive.ObjectIDFromHex(webpageId)

	err := repository.UpdateIssue(issueUpdateBody, wpId, ishId).Decode(&issue)


	if err != nil {
		return nil, err

	}
	return issueUpdateBody, nil
}
