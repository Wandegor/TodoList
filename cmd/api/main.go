package main

import (
	"ISpringTODOList/internal/database"
	"ISpringTODOList/internal/handlers"
	"ISpringTODOList/internal/models"
	"fmt"
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

	tasks := []models.Task{
		{
			Text: "first task",
		},
		{
			Text: "another",
		},
	}

	for _, task := range tasks {
		if err := db.Create(&task).Error; err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskHandler.GetTasks(w)
		}
		if r.Method == http.MethodPost {
			taskHandler.CreateTask(w, r)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
