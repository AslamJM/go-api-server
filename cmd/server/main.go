package main

import (
	"kstrategy-backend/config"
	"kstrategy-backend/internals/api/handlers"
	"kstrategy-backend/internals/db"
	"kstrategy-backend/internals/models"
	"kstrategy-backend/internals/repositories"
	"kstrategy-backend/internals/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	DB := db.InitDB()
	r := mux.NewRouter()
	models.NewValidator()

	userRepo := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(*userService)

	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
