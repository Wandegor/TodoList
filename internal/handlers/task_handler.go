package handlers

import (
	"ISpringTODOList/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{
		db: db,
	}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter) {
	var tasks []models.Task

	if err := h.db.Find(&tasks).Error; err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println(err)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.db.Create(&task).Error; err != nil {
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		return
	}
	return
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString) //конверт в int
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	if err := h.db.First(&task, id).Error; err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	if err := h.db.Delete(&task).Error; err != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
