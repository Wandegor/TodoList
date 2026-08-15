package handlers

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		http.Error(w, appErrors.ErrGetTasks.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println(err)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, appErrors.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateTask(&task)
	if err != nil {
		// TODO: придумать как разделить ошибку на техническую/пользовательскую
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println(err)
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idString, 10, 64) //конверт в uint
	if err != nil {
		http.Error(w, appErrors.ErrInvalidTaskID.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(uint(id))
	if err != nil {
		if errors.Is(err, appErrors.ErrTaskNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		log.Println(err)
		http.Error(w, appErrors.ErrInternalServerError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
