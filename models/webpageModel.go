package models

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/utilities"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveWebpageScan(wp *entity.Webpage) (*entity.Webpage, error) {

	if wp.Issue == nil {
		wp.Issue = make([]entity.Issue, 0) // this is alloc free
	}else {
		for _, s:= range wp.Issue{
			s.Timestamp = time.Now().Format("2006-01-02 15:04:05")
			s.ID = primitive.NewObjectID()
			fmt.Println("issue here", s)

		}
	}
	fmt.Println("out here", wp.Issue)



	result, err := database.WebpageCollection.InsertOne(database.Ctx, wp)
	if err != nil {
		fmt.Println("unable to insert record", err)
		return nil, err
	}
	wp.ID = result.InsertedID.(primitive.ObjectID)
	if wp.Website.ID == primitive.NilObjectID {
		web, err := GetWebsiteByField(wp.Website.Url)
		fmt.Println("see web", web)
		if web == nil {
			_, err := database.WebsiteCollection.InsertOne(database.Ctx, wp.Website)
			if err != nil {
				fmt.Println("unable to insert website", err)
			}
			fmt.Println("created a new website", err)
		}
		fmt.Println("website already exist", err)

	}
	return GetWebpageById(result.InsertedID.(primitive.ObjectID))

}

func GetAllWebpages() ([]entity.Webpage, error) {
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

func GetWebpageById(id primitive.ObjectID) (*entity.Webpage, error) {
	var wp entity.Webpage
	
	err := database.WebpageCollection.
		FindOne(database.Ctx, bson.D{{"_id", id}}).
		Decode(&wp)
	if err != nil {
		return nil, err
	}
	return &wp, nil
}

func GetWebpageByField(field string) ([]entity.Webpage, error) {
	var wp entity.Webpage
	var webpages []entity.Webpage

	cursor, err := database.WebpageCollection.
		Find(database.Ctx,
			bson.M{
				"$or": bson.A{
					bson.M{
						"name": &field,
					},
					bson.M{
						"website.name": &field,
					},
					bson.M{
						"url": &field,
					},
					bson.M{
						"website.url": &field,
					},
				},
			})

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

func DeleteWebpage(id primitive.ObjectID) error {
	_, err := database.WebpageCollection.DeleteOne(database.Ctx, bson.D{{"_id", id}})
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
	v, err = GetWebpageById(utilities.StringToPrimitive(id))
	if err != nil {
		return nil, err
	}
	return v, err
}
