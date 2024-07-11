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

var _ = godotenv.Load()

func init() {
	connectionString := os.Getenv("CONNECT_STRING")
	db := "devblogs"
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected To Database\n")
	BlogCollection := client.Database(db).Collection("Blogs")
	fmt.Println(BlogCollection)

}
