package main

import (
	"Book-Management/pkg/config"
	"Book-Management/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//  Initialize database connection before handling requests
	config.Connect()
	// config.ExistsDatabase()

	//  Ensure database instance is accessible
	db := config.GetDB()
	if db == nil {
		log.Fatal("Database connection failed, ensure Connect() was called.")
	}

	//  Set up router
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Println("Server running on localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
