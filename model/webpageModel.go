package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SaveWebpageScan(data *dto.WebpageRequestBody) (*entity.Webpage, error) {



	if data.Issue == nil {
		data.Issue = make([]dto.IssueRequestBody, 0) // this is alloc free
	}
	
	// else {
	// 	for _, s:= range wp.Issue{
	// 		s.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	// 		s.ID = primitive.NewObjectID()

	// 	}
	// }
	wpage := data.ToWebpageEntities()



	result, err := repository.SaveWebpage(wpage)
	if err != nil {
		return nil, err
	}
	wpage.ID = result.InsertedID.(primitive.ObjectID)
	if wpage.Website.ID == primitive.NilObjectID {
		web, err := GetWebpageByField(wpage.Website.Url)
		if web == nil {
			_, err := database.WebsiteCollection.InsertOne(database.Ctx, wpage.Website)
			if err != nil {
				fmt.Println("unable to insert website", err)
			}
			fmt.Println("created a new website", err)
		}
		fmt.Println("website already exist", err)

	}
	return wpage,nil

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

func GetWebpageById(id string) (*entity.Webpage, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	wp, err := repository.FindWebpageById(objectId)
	if err != nil {
		return nil, err
	}
	return &wp, nil
}

func GetWebpageByField(field string) ([]entity.Webpage, error) {
	var wp entity.Webpage
	var webpages []entity.Webpage

	cursor, err := repository.GetWebpageByField(field)

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

func DeleteWebpage(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteWebpage(objectId)
	if err != nil {
		return err
	}
	return err
}

func UpdateWebpage(v *entity.Webpage, id string) (*entity.Webpage, error) {

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
