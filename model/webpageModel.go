package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveWebpageScan(data *dto.WebpageRequestBody) (dto.WebpageResponseBody, error) {
	var resp dto.WebpageResponseBody

	if data.Issue == nil {
		data.Issue = make([]dto.IssueRequestBody, 0) // this is alloc free
	}

	webpageEntity := data.ToWebpageEntities()

	website, err := GetWebsiteByField(webpageEntity.Website.Name)
	var ws *entity.Website

	//if website doesnt exist , create a new website for that webpage
	if website == nil {
		ws, err = CreateWebsite(&data.Website)
		if err != nil {
			fmt.Println("unable to insert website", err)
		} else {
			fmt.Println("created a new website", err)
			webpageEntity.Website = *ws
		}
	} else {
		//website exists so dont create a new website
		webpageEntity.Website = *website
	}

	_, err = repository.SaveWebpage(webpageEntity)
	if err != nil {
		return resp, err
	}
	resp.ID = webpageEntity.ID.Hex()

	return resp, nil

}

func GetAllWebpages() ([]dto.WebpageResponseBody, error) {
	var webpage entity.Webpage
	var webpageResponse dto.WebpageResponseBody
	var webpages []dto.WebpageResponseBody
	cursor, err := repository.FindWebpages()
	if err != nil {
		defer cursor.Close(database.Ctx)
		return webpages, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&webpage)
		if err != nil {
			return webpages, err
		}

		webpageResponse.ID = webpage.ID.Hex()
		webpageResponse.Name = webpage.Name
		webpageResponse.ScanTime = webpage.ScanTime
		webpageResponse.Url = webpage.Url
		webpageResponse.Website = webpage.Website.Name
		webpages = append(webpages, webpageResponse)
	}
	return webpages, nil
}

func GetWebpageById(objectId primitive.ObjectID) (dto.WebpageFullResponseBodyNew, error) {
	var resp dto.WebpageFullResponseBodyNew
	
	wp, err := repository.FindWebpageById(objectId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return resp, errors.New("no webpage with the provided id")
		}
		return resp, err
	}

	// resp.FoundStats = getFoundtypeStats(wp)
	// resp.ImpactStats = getImpactStats(wp)

	resp.FoundStats = getFoundStatNew(wp)
	resp.ImpactStats = getImpactStatNew(wp)

	resp.ID = wp.ID.Hex()
	resp.Issue = wp.Issue
	resp.Name = wp.Name
	resp.Note = wp.Note
	resp.Website = wp.Website
	resp.ScanTime = wp.ScanTime
	resp.Url = wp.Url

	return resp, nil
}

func GetWebpageByField(searchField string, sortByField string, orderBy int, pageSize int64, pageNum int64) (dto.WebpagesResult, error) {
	var webpage entity.Webpage
	var webpageResponse dto.WebpageResponseBody
	var webpages []dto.WebpageResponseBody
	var webpagesResult dto.WebpagesResult

	cursor, err, total := repository.GetWebpageByField(searchField, sortByField, orderBy, pageSize, pageNum)

	if err != nil {
		defer cursor.Close(database.Ctx)
		return webpagesResult, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&webpage)
		if err != nil {
			return webpagesResult, err
		}
		webpageResponse.ID = webpage.ID.Hex()
		webpageResponse.Name = webpage.Name
		webpageResponse.ScanTime = webpage.ScanTime
		webpageResponse.Url = webpage.Url
		webpageResponse.Website = webpage.Website.Name
		webpages = append(webpages, webpageResponse)
	}

	webpagesResult.Data = webpages
	webpagesResult.TotalCount = total
	webpagesResult.Page = int(pageNum)
	webpagesResult.LastPage = int(math.Ceil(float64(total)/float64(pageSize))) 

	return webpagesResult, nil
}



func DeleteWebpage(objectId primitive.ObjectID) error {
	cursor, err := repository.DeleteWebpage(objectId)
	if(cursor.DeletedCount ==0 ){
		return errors.New("No webpage with the provided id")
	}
	if err != nil {
		return err
	}
	log.Print("del", err)
	return err
}

// func UpdateWebpage(v *entity.Webpage, id string) (*entity.Webpage, error) {

// 	result, err := database.WebpageCollection.UpdateOne(database.Ctx, bson.M{"_id": id},
// 		bson.M{
// 			"$set": &v,
// 		},
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if result.MatchedCount == 0 {
// 		return nil, err
// 	}
// 	if result.ModifiedCount == 0 {
// 		return nil, err
// 	}
// 	v, err = GetWebpageById(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return v, err
// }


func getImpactStatNew(wp entity.Webpage) []dto.ImpactStatNew {

	var aggResult = make([]dto.ImpactStatNew, 0)
	var minorCount int = 0
	var moderateCount int = 0
	var seriousCount int = 0
	var criticalCount int = 0

	for _, iss := range wp.Issue {
		switch strings.ToLower(iss.Impact) {
		case "serious":
			seriousCount += len(iss.Occurence)
		case "minor":
			minorCount += len(iss.Occurence)
		case "moderate":
			moderateCount += len(iss.Occurence)
		case "critical":
			criticalCount += len(iss.Occurence)
		}
	}
    seriousObject := dto.ImpactStatNew{Impact: "serious", Count: seriousCount}
    minorObject := dto.ImpactStatNew{Impact: "minor", Count: minorCount}
    moderateObject := dto.ImpactStatNew{Impact: "moderate", Count: moderateCount}
    criticalObject := dto.ImpactStatNew{Impact: "critical", Count: criticalCount}

	aggResult =	append(aggResult, seriousObject)
	aggResult =	append(aggResult, minorObject)
	aggResult =	append(aggResult, moderateObject)
	aggResult =	append(aggResult, criticalObject)

	return aggResult
}

func getFoundStatNew(wp entity.Webpage) []dto.FoundStatNew {

	var aggResult = make([]dto.FoundStatNew, 0)
	var automaticCount int = 0
	var guidedCount int = 0
	var manualCount int = 0
	var needsReviewCount int = 0

	for _, iss := range wp.Issue {
		switch strings.ToLower(iss.Found) {
		case "automatic":
			automaticCount += len(iss.Occurence)
		case "guided":
			guidedCount += len(iss.Occurence)
		case "manual":
			manualCount += len(iss.Occurence)
		}
		for _, occ := range iss.Occurence {
			if(occ.NeedsReview){
				needsReviewCount ++
			}
		}
	}
    automaticObject := dto.FoundStatNew{Found: "automatic", Count: automaticCount}
    guidedObject := dto.FoundStatNew{Found: "guided", Count: guidedCount}
    needsReviewObject := dto.FoundStatNew{Found: "needsReview", Count: needsReviewCount}
    manualObject := dto.FoundStatNew{Found: "manual", Count: manualCount}

	aggResult =	append(aggResult, automaticObject)
	aggResult =	append(aggResult, guidedObject)
	aggResult =	append(aggResult, needsReviewObject)
	aggResult =	append(aggResult, manualObject)

	return aggResult
}