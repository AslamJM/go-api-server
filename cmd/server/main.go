package main

import (
	"kstrategy-backend/config"
	"kstrategy-backend/internals/api/handlers"
	"kstrategy-backend/internals/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
