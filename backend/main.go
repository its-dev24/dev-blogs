package main

import (
	"log"
	"net/http"
	"os"

	"github.com/its-dev24/dev-blogs/router"
	"github.com/joho/godotenv"
)

func main() {

	mux := router.Router()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while getting Env variable : ", err)
	}
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Coudn't start server : ", err)
	}
}
