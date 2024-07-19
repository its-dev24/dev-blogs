package database

import (
	"context"
	"fmt"
	"log"

	"github.com/its-dev24/dev-blogs/modals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneBlog(blog modals.Blog) (string, error) {
	result, err := BlogCollection.InsertOne(context.Background(), blog)
	if err != nil {
		fmt.Println("Erro While Inserting blog : ", err)
		return "", err
	}
	fmt.Printf("ID of Rows inserted : %v", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func UpdateBlog(id string, blogs modals.Blog) (int, error) {
	blogObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Error While convert blogID(update Blog)", err)
	}
	filer := bson.M{"_id": blogObjectID}
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
	blogIdObject, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		log.Fatal("error while creating id : ", err)
	}
	filter := bson.M{"_id": blogIdObject}
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
