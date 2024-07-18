package router

import (
	"net/http"

	"github.com/its-dev24/dev-blogs/controllers"
)

func Router() *http.ServeMux {
	var mux http.ServeMux
	mux.HandleFunc("GET /", controllers.HomeScreen)
	mux.HandleFunc("GET /api/blogs", controllers.FetchAllBlogs)
	mux.HandleFunc("DELETE /api/blogs", controllers.DeleteAllBlogs)
	mux.HandleFunc("PUT /api/blogs/{id}", controllers.UpdateBlog)
	return &mux
}
