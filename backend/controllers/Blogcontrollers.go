package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "github.com/its-dev24/dev-blogs/Database"
)

//Fetch all Blogs

func FetchAllBlogs(w http.ResponseWriter, r *http.Request) {
	result := database.FindAllBlog()
	if result.Error != nil {
		json.NewEncoder(w).Encode("Error while getting Result")
	} else if result.Value == nil {
		json.NewEncoder(w).Encode("No Blog Data")
	}
	json.NewEncoder(w).Encode(result.Value)
}

//Delete All values

func DeleteAllBlogs(w http.ResponseWriter, r *http.Request) {

	deleteCount, err := database.DeleteAllBlogs()
	if err != nil {
		json.NewEncoder(w).Encode("Error while deleting Blogs " + err.Error())
		return
	}
	json.NewEncoder(w).Encode("No  of Items Deleted" + strconv.Itoa(deleteCount))

}

//