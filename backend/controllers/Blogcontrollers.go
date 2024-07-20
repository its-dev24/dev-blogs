package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/its-dev24/dev-blogs/Database"
	"github.com/its-dev24/dev-blogs/helper"
	"github.com/its-dev24/dev-blogs/modals"
)

//Home screen

func HomeScreen(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>HOME SCREEN</h1>"))
}

//Fetch all Blogs

func FetchAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := database.FindAllBlog()
	if result.Error != nil {
		msg := helper.CreateMap("error", "Error while Getting Result"+result.Error.Error())
		json.NewEncoder(w).Encode(msg)
	} else if result.Value == nil {
		msg := helper.CreateMap("msg", "No Blog Data")
		json.NewEncoder(w).Encode(msg)
		return
	}
	json.NewEncoder(w).Encode(result.Value)
	fmt.Println(result.Value)
}

//Delete All values

func DeleteAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deleteCount, err := database.DeleteAllBlogs()
	if err != nil {
		json.NewEncoder(w).Encode("Error while deleting Blogs " + err.Error())
		return
	}
	msg := helper.CreateMap("msg", "No of Blogs Deleted"+strconv.Itoa(deleteCount))
	json.NewEncoder(w).Encode(msg)

}

//Delete a Single Blog

func DeleteABlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	deleteCount := database.DeleteBlog(id)
	msg := helper.CreateMap("msg", "No of Blogs Deleted"+strconv.Itoa(deleteCount))
	json.NewEncoder(w).Encode(msg)
}

// Update a Blog
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var blogData modals.Blog
	id := r.PathValue("id")
	json.NewDecoder(r.Body).Decode(&blogData)
	updateCount, err := database.UpdateBlog(id, blogData)
	if err != nil {
		json.NewEncoder(w).Encode("Error While Inserting : " + err.Error())
		return
	}
	msg := helper.CreateMap("msg", "No of Blogs Updated"+strconv.Itoa(updateCount))

	json.NewEncoder(w).Encode(msg)

}

//Insert  A Blog

func InsertABlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
