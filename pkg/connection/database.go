package connection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"

	"file-system/pkg/helpers"
)

type Connection interface {
	MongoDbConnection() *mongo.Collection
}

type Collections struct {
	File *mongo.Collection
	User *mongo.Collection
}

// func DatabaseConnection() *gorm.DB {
// 	dsn := "host=localhost user=postgres password=password dbname=godb port=5432 sslmode=disable TimeZone=Asia/Tashkent"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	helpers.ErrorHelper(err)

// 	return db
// }

func DatabaseConnection() Collections {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	helpers.ErrorHelper(err)

	database := client.Database("file-system")
	fileCollection := database.Collection("files")
	userCollection := database.Collection("users")

	collections := Collections{
		File: fileCollection,
		User: userCollection,
	}

	return collections
}
