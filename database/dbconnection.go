package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	CriteriaCollection *mongo.Collection
	WebpageCollection  *mongo.Collection
	WebsiteCollection  *mongo.Collection
	Ctx                = context.TODO()
)

func Setup() {
	host := "127.0.0.1"
	port := "27017"
	connectionUri := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionUri)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("AssessibilityDB")
	CriteriaCollection = db.Collection("criteria")
	WebpageCollection = db.Collection("webpage")
	WebsiteCollection = db.Collection("website")
}
