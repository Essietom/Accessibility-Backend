package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CriteriaCollection *mongo.Collection
	WebpageCollection  *mongo.Collection
	WebsiteCollection  *mongo.Collection
	Ctx                = context.TODO()
)

func Setup() {
	connectionUri := os.Getenv("DB_CONNECTION_URI")

	clientOptions := options.Client().ApplyURI(connectionUri)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(os.Getenv("DB_NAME"))
	CriteriaCollection = db.Collection("criteria")
	WebpageCollection = db.Collection("webpage")
	WebsiteCollection = db.Collection("website")
}
