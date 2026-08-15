package main

import (
	"ISpringTODOList/internal/database"
	"ISpringTODOList/internal/handlers"
	"ISpringTODOList/internal/models"
	"log"
	"net/http"
)

func main() {

	log.Printf("Server Start on port 8080")
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	taskHandler := handlers.NewTaskHandler(db)
	log.Printf("Database Connect Success")

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database AutoMigrate Success")

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskHandler.GetTasks(w)
		}
		if r.Method == http.MethodPost {
			taskHandler.CreateTask(w, r)
		}
		if r.Method == http.MethodDelete {
			taskHandler.DeleteTask(w, r)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
