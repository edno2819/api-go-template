package database

import (
	"context"
	"gin-mongo-api/setting"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo_url := setting.DatabaseSetting.Url
	log.Println("mongo_url", mongo_url)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	database_name := setting.DatabaseSetting.Database
	database := client.Database(database_name)
	return database
}
