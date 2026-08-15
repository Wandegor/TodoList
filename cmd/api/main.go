package main

import (
	"ISpringTODOList/internal/database"
	"ISpringTODOList/internal/models"
	"encoding/json"
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
		var tasks []models.Task

		if err := db.Find(&tasks).Error; err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			log.Println(err)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
