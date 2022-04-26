package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SaveWebpageScan(data *dto.WebpageRequestBody) (*entity.Webpage, error) {

	if data.Issue == nil {
		data.Issue = make([]dto.IssueRequestBody, 0) // this is alloc free
	}
	
	webpageEntity := data.ToWebpageEntities()

	website, err := GetWebsiteByField(webpageEntity.Website.Name)
	var ws *entity.Website

	if website == nil {
		ws, err = CreateWebsite(&data.Website)
		if err != nil {
			fmt.Println("unable to insert website", err)
		}else{
		fmt.Println("created a new website", err)
		webpageEntity.Website = *ws
		}
	}else{
	fmt.Println("website already exist", err)
	webpageEntity.Website = *website
	}


	_, err = repository.SaveWebpage(webpageEntity)
	if err != nil {
		return nil, err
	}

	return webpageEntity,nil

}

func GetAllWebpages() ([]entity.Webpage, error) {
	var webpage entity.Webpage
	var webpages []entity.Webpage
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
		webpages = append(webpages, webpage)
	}
	return webpages, nil
}

func GetWebpageById(id string) (dto.WebpageFullResponseBody, error) {
	var resp dto.WebpageFullResponseBody
	objectId, _ := primitive.ObjectIDFromHex(id)

	wp, err := repository.FindWebpageById(objectId)
	if err != nil {
		//return nil, err
	}
	resp.FoundStats = getFoundtypeStats(wp)
	resp.ImpactStats = getImpactStats(wp)
	resp.ID = string(wp.ID.String())
	resp.Issue = wp.Issue
	resp.Name = wp.Name
	resp.Note = wp.Note
	resp.Website = wp.Website
	resp.ScanTime = wp.ScanTime
	resp.Url = wp.Url

	return resp, nil
}

func GetWebpageByField(searchField string, sortByField string, orderBy string, pageSize int64, pageNum int64) ([]entity.Webpage, error) {
	var wp entity.Webpage
	var webpages []entity.Webpage
	var order int 

	if orderBy == "asc" {
		order = 1
	}else{
		order = -1
	}


	cursor, err := repository.GetWebpageByField(searchField, sortByField, order, pageSize , pageNum )

	if err != nil {
		defer cursor.Close(database.Ctx)
		return webpages, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&wp)
		if err != nil {
			return webpages, err
		}
		webpages = append(webpages, wp)
	}
	return webpages, nil
}

func getImpactStats(wp entity.Webpage)dto.ImpactStat{

	var minorCount int = 0
	var moderateCount int = 0
	var seriousCount int = 0
	var criticalCount int = 0
	var totalCount int = 0
	var aggResult dto.ImpactStat

	for _, iss := range wp.Issue{
		switch iss.Impact{
		case "Serious":
			seriousCount+=1
		case "Minor":
			minorCount+=1
		case "Moderate":
			moderateCount+=1
		case "Critical":
			criticalCount+=1
		}
	}
	aggResult.Critical = criticalCount
	aggResult.Moderate = moderateCount
	aggResult.Minor = minorCount
	aggResult.Serious = seriousCount
	aggResult.ImpactTotal = totalCount

	return aggResult
}

func getFoundtypeStats(wp entity.Webpage)dto.FoundStat{

	var automaticCount int = 0
	var guidedCount int = 0
	var needsReviewCount int = 0
	var totalCount int = 0
	var aggResult dto.FoundStat

	for _, iss := range wp.Issue{
		switch iss.Impact{
		case "Automatic":
			automaticCount+=1
		case "Guided":
			guidedCount+=1
		case "NeedsReview":
			needsReviewCount+=1
		}
	}
	aggResult.Automatic = automaticCount
	aggResult.Guided = guidedCount
	aggResult.NeedsReview = needsReviewCount
	aggResult.FoundTotal = totalCount
	return aggResult
}

func DeleteWebpage(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteWebpage(objectId)
	if err != nil {
		return err
	}
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
