package database

import (
	"Accessibility-Backend/utilities"
	"context"
	"log"
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
	// connectionUri := os.Getenv("DB_CONNECTION_URI")

	config, err := utilities.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }
	connectionUri := config.DBConnectionUri

	clientOptions := options.Client().ApplyURI(connectionUri)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.DBName)
	CriteriaCollection = db.Collection("criteria")
	WebpageCollection = db.Collection("webpage")
	WebsiteCollection = db.Collection("website")
}
