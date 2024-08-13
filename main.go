package main

import (
	"log"
	"net/http"

	"mongo-crud/conn"
	"mongo-crud/controller"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize MongoDB connection
	client := conn.Db()

	// Set the collection
	controller.UserCollection = client.Database("mongo_crud").Collection("users")

	router := mux.NewRouter()

	// Route Handlers
	router.HandleFunc("/create", controller.CreateProfile).Methods("POST")
	router.HandleFunc("/user", controller.GetUserProfile).Methods("POST")
	router.HandleFunc("/update", controller.UpdateProfile).Methods("PUT")
	router.HandleFunc("/delete/{id}", controller.DeleteProfile).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
