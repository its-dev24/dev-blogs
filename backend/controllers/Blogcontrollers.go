package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/its-dev24/dev-blogs/Database"
	"github.com/its-dev24/dev-blogs/modals"
)

//Home screen

func HomeScreen(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("<h1>HOME SCREEN</h1>")
}

//Fetch all Blogs

func FetchAllBlogs(w http.ResponseWriter, r *http.Request) {
	result := database.FindAllBlog()
	if result.Error != nil {
		json.NewEncoder(w).Encode("Error while getting Result")
	} else if result.Value == nil {
		json.NewEncoder(w).Encode("No Blog Data")
		return
	}
	json.NewEncoder(w).Encode(result.Value)
	fmt.Println(result.Value)
}

//Delete All values

func DeleteAllBlogs(w http.ResponseWriter, r *http.Request) {

	deleteCount, err := database.DeleteAllBlogs()
	if err != nil {
		json.NewEncoder(w).Encode("Error while deleting Blogs " + err.Error())
		return
	}
	json.NewEncoder(w).Encode("No  of Items Deleted : " + strconv.Itoa(deleteCount))

}

//Delete a Single Blog

func DeleteABlog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	deleteCount := database.DeleteBlog(id)
	json.NewEncoder(w).Encode("No of Blogs Deleted : " + string(deleteCount))
}

// Update a Blog
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	var blogData modals.Blog
	id := r.PathValue("id")
	json.NewDecoder(r.Body).Decode(&blogData)
	updateCount, err := database.UpdateBlog(id, blogData)
	if err != nil {
		json.NewEncoder(w).Encode("Error While Inserting : " + err.Error())
		return
	}
	json.NewEncoder(w).Encode("No of Filed Updated : " + strconv.Itoa(updateCount))

}

//Insert  A Blog

func InsertABlog(w http.ResponseWriter, r *http.Request) {
	var blogData modals.Blog
	json.NewDecoder(r.Body).Decode(&blogData)
	if blogData.IsEmpty() {
		json.NewEncoder(w).Encode("Empty Json Body")
		fmt.Println("Json body is empty..")
		return
	}
	insertId, err := database.InsertOneBlog(blogData)
	if err != nil {
		json.NewEncoder(w).Encode("Error while inserting Value : " + insertId)
		return
	}
	json.NewEncoder(w).Encode("Inserted Succesfully with id : " + insertId)
}
