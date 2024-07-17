package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var BlogCollection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Coudn't load env variables", err)
	}
	connectionString := os.Getenv("CONNECT_STRING")
	db := "devblogs"
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Could not create client ", err)
	}
	fmt.Printf("Connected To Database\n")
	BlogCollection = client.Database(db).Collection("Blogs")
	if BlogCollection != nil {
		fmt.Println("Collection created")
	}
	// fmt.Println(BlogCollection)

}
