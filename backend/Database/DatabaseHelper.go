package database

import (
	"context"
	"fmt"
	"log"

	"github.com/its-dev24/dev-blogs/helper"
	"github.com/its-dev24/dev-blogs/modals"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertOneBlog(blog modals.Blog) {
	result, err := BlogCollection.InsertOne(context.Background(), blog)
	helper.CheckError(err)
	fmt.Printf("ID of Rows inserted : %v", result.InsertedID)
}

func UpdateBlog(blogs modals.Blog) (int, error) {
	filer := bson.M{"_id": blogs.Id}
	update := bson.M{"$set": bson.M{"title": blogs.Title, "author": blogs.Author, "body": blogs.BlogBody}}
	result, err := BlogCollection.UpdateOne(context.Background(), filer, update)
	if err != nil {
		log.Fatal("Error during Insertion : ", err)
		return 0, err

	}
	fmt.Printf("Update Sucessfully\n")
	fmt.Printf("No of row affected : %v\n", result.MatchedCount)
	return int(result.MatchedCount), nil
}

func DeleteBlog(blogId string) int {
	filter := bson.M{"_id": blogId}
	result, err := BlogCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("Error While Deleting blog : ", err)
	}
	fmt.Println("Deleted Succesfully")
	fmt.Printf("No of documents affected : %v\n", result.DeletedCount)
	return int(result.DeletedCount)
}

func DeleteAllBlogs() (int, error) {
	result, err := BlogCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	fmt.Println("Deleted Succesfully")
	fmt.Println("No of documents affected : ", result.DeletedCount)
	return int(result.DeletedCount), nil
}

func FindAllBlog() modals.ReturnValue {
	cursor, err := BlogCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("Error while Fetching Blogs.. : ", err)
	}
	var blogs []bson.M

	err = cursor.All(context.Background(), &blogs)
	returnValue := modals.ReturnValue{Error: err, Value: blogs}
	fmt.Println("All blogs Fetched!!..")
	return returnValue
}
