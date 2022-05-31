package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//update website, add issue to issue object
func AddIssue(data dto.IssueRequestBody, websiteId string) (*entity.Issue, error) {

	if data.Criteria == nil {
		data.Criteria = make([]dto.CriteriaRequestBody, 0) // this is alloc free
	}
	if data.Occurence == nil {
		data.Occurence = make([]dto.OccurenceRequestBody, 0) // this is alloc free
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

func GetOccurenceById(issueId string, webpageId string, occurenceId string) (*entity.Occurence, error) {
	var occurence entity.Occurence
	ishId, _ := primitive.ObjectIDFromHex(issueId)	
	wpId, _ := primitive.ObjectIDFromHex(webpageId)
	occId, _ := primitive.ObjectIDFromHex(occurenceId)


	cursor, err := repository.FindOccurenceById(ishId, wpId, occId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("occurence not found")
		}
		return nil, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&occurence)
		if err != nil {
			log.Println("some error occured while decoding")
			return &occurence, err

		}
	}
	return &occurence, nil
}

func GetOccurenceCountForIssue(issueId string, webpageId string) (int){
	var cnt entity.Count
	ishId, _ := primitive.ObjectIDFromHex(issueId)	
	wpId, _ := primitive.ObjectIDFromHex(webpageId)
	cursor, err := repository.GetOccurenceCount(ishId, wpId)
	if err != nil {
		return 0
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&cnt)
		if err != nil {
		return 0
		}
	}
	return int(cnt.Count)
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

//updates an issue under a website and the occurences under that issue
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

func UpdateOccurence(occurenceUpdateBody *entity.Occurence, webpageId string, issueId string, occurenceId string) (*entity.Occurence, error) {

	// var occurence entity.Occurence
	var webpage entity.Webpage

	ishId, _ := primitive.ObjectIDFromHex(issueId)	
	wpId, _ := primitive.ObjectIDFromHex(webpageId)
	occId, _ := primitive.ObjectIDFromHex(occurenceId)

	err := repository.UpdateOccurence(occurenceUpdateBody, wpId, ishId, occId).Decode(&webpage)
	fmt.Println("see web occ", webpage.Issue)
	if err != nil {
		return nil, err

	}
	return occurenceUpdateBody, nil
}




//update website websiteid, pull issue where webpage id = webpage and issue id = issueid
func DeleteOccurence(webpageId string, issueId string, occurenceId string) error {

	ishId, _ := primitive.ObjectIDFromHex(issueId)	
	wpId, _ := primitive.ObjectIDFromHex(webpageId)
	occId, _ := primitive.ObjectIDFromHex(occurenceId)
	
	result, err := repository.DeleteOccurence(wpId, ishId, occId)
	if err != nil {
		return errors.New("some error occurred")
	}
	if result.MatchedCount == 0 {
		return  errors.New("no webpage found with the id provided")
	}
	// if result.ModifiedCount == 0 {
	// 	log.Print("can't delete occurence because it doesnt exist");
	// 	return errors.New("the occurence id provided does  not exist for the provided issue id,occurrence was not successfully deleted")
	// }
	log.Print("An issue occurence got deleted");

	return  nil

	}

	

	
func DeleteOccurenceOrIssue(webpageId string, issueId string, occurenceId string) error{
		
	occurenceCount := GetOccurenceCountForIssue(issueId,webpageId)

	if occurenceCount <= 1{
		return DeleteIssue(webpageId, issueId)
	}else{
		return DeleteOccurence(webpageId, issueId, occurenceId)
	}

}