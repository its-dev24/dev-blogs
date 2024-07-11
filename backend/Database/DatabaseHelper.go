package database

import (
	"context"
	"fmt"
	"log"

	"github.com/its-dev24/dev-blogs/helper"
	"github.com/its-dev24/dev-blogs/modals"
	"go.mongodb.org/mongo-driver/bson"
)

func insertOneBlog(blog modals.Blog) {
	result, err := BlogCollection.InsertOne(context.Background(), blog)
	helper.CheckError(err)
	fmt.Printf("ID of Rows inserted : %v", result.InsertedID)
}

func updateBlog(blogs modals.Blog) {
	filer := bson.M{"_id": blogs.Id}
	update := bson.M{"$set": bson.M{"title": blogs.Title, "author": blogs.Author, "body": blogs.BlogBody}}
	result, err := BlogCollection.UpdateOne(context.Background(), filer, update)
	if err != nil {
		log.Fatal("Error during Insertion : ", err)
	}
	fmt.Printf("Inserted Sucessfully\n")
	fmt.Printf("No of row affected : %v\n", result.MatchedCount)
}

func deleteBlog(blogs modals.Blog) {
	filter := bson.M{"_id": blogs.Id}
	result, err := BlogCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("Error While Deleting blog : ", err)
	}
	fmt.Println("Deleted Succesfully")
	fmt.Printf("No of documents affected : %v\n", result.DeletedCount)
}

func deleteAllBlogs() {
	result, err := BlogCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("Error while deleting all blogs : ", err)
	}
	fmt.Println("Deleted Succesfully")
	fmt.Println("No of documents affected : ", result.DeletedCount)
}
