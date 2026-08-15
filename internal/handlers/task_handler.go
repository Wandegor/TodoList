package handlers

import (
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

//func (h *TaskHandler) GetTasks(w http.ResponseWriter) {
//	var tasks []models.Task
//
//	if err := h.db.Find(&tasks).Error; err != nil {
//		log.Fatal(err)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//
//	if err := json.NewEncoder(w).Encode(tasks); err != nil {
//		log.Println(err)
//	}
//}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err != json.NewEncoder(w).Encode(task) {
		log.Println(err)
	}
}

//func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
//	var task models.Task
//
//	idString := r.URL.Query().Get("id")
//
//	id, err := strconv.Atoi(idString) //конверт в int
//	if err != nil {
//		http.Error(w, "invalid task id", http.StatusBadRequest)
//		return
//	}
//
//	if err := h.db.First(&task, id).Error; err != nil {
//		http.Error(w, "task not found", http.StatusNotFound)
//		return
//	}
//
//	if err := h.db.Delete(&task).Error; err != nil {
//		http.Error(w, "failed to delete task", http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
