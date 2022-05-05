package model

import (
	"Accessibility-Backend/database"
	"Accessibility-Backend/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCriteria(v *entity.Criteria) (*entity.Criteria, error) {
	_, err := database.CriteriaCollection.InsertOne(database.Ctx, v)
	if err != nil {
		fmt.Println("unable to insert record", err)
		return nil, err
	}
	//v.ID = result.InsertedID.(primitive.ObjectID)
	return v, err
}

func GetAllCriteria() ([]entity.Criteria, error) {
	var criterion entity.Criteria
	var criteria []entity.Criteria
	cursor, err := database.CriteriaCollection.Find(database.Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(database.Ctx)
		return criteria, err
	}

	for cursor.Next(database.Ctx) {
		err := cursor.Decode(&criterion)
		if err != nil {
			return criteria, err
		}
		criteria = append(criteria, criterion)
	}
	return criteria, nil
}

func GetCriteriaById(id string) (*entity.Criteria, error) {
	var v entity.Criteria
	err := database.CriteriaCollection.
		FindOne(database.Ctx, bson.D{{"_id", id}}).
		Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func DeleteCriteria(id primitive.ObjectID) error {
	_, err := database.CriteriaCollection.DeleteOne(database.Ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return err
}

func UpdateCriteria(v *entity.Criteria, id string) (*entity.Criteria, error) {

	result, err := database.CriteriaCollection.UpdateOne(database.Ctx, bson.M{"_id": id},
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
	v, err = GetCriteriaById(id)
	if err != nil {
		return nil, err
	}
	return v, err
}
