package main

import (
	"ISpringTODOList/internal/config"
	"ISpringTODOList/internal/database"
	"ISpringTODOList/internal/handlers"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
	"ISpringTODOList/internal/router"
	"ISpringTODOList/internal/services"
	"log"
	"net/http"
	"os"
)

func main() {
	mainDbConfig := config.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
	db, err := database.Connect(mainDbConfig)
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewTaskRepository(db)
	service := services.NewTaskService(repository)
	taskHandler := handlers.NewTaskHandler(service)
	log.Printf("Database Connect Success")

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database AutoMigrate Success")

	mux := router.SetupRouter(taskHandler)

	log.Printf("Server Start on port 8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
