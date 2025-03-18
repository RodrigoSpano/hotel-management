package db_mongo

import (
	"fmt"

	"github.com/gofiber/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectDB() *mongo.Client {
	store := mongodb.New(mongodb.Config{
		ConnectionURI: "mongodb://user:password@127.0.0.1:27017",
		Reset:         false,
	})
	fmt.Println("Connected to MongoDB\n")
	return store.Conn().Client()
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("hotel-manager").Collection(collectionName)
	return collection
}

var DB *mongo.Client = ConnectDB()
