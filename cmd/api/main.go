package main

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/database"
	"ISpringTODOList/internal/handlers"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
	"ISpringTODOList/internal/services"
	"log"
	"net/http"
)

func main() {

	db, err := database.Connect()
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

	http.HandleFunc("/tasks/archive", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, appErrors.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
			return
		}

		taskHandler.GetArchivedTasks(w)
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTasks(w)
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		default:
			http.Error(w, appErrors.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		}

	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)
		case http.MethodPatch:
			taskHandler.CompleteTask(w, r)
		default:
			http.Error(w, appErrors.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		}

	})

	log.Printf("Server Start on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
