package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nurhidaylma/golang-simple-login/config"
	"github.com/nurhidaylma/golang-simple-login/service"
)

func main() {
	config.InitialMigration()

	router := *mux.NewRouter()

	router.HandleFunc("/register", service.Register).Methods("POST")
	router.HandleFunc("/login", service.Login).Methods("POST")

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(&router))
	if err != nil {
		log.Fatal(err)
	}
}
