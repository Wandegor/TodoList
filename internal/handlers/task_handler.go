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
	"strings"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) writeTasks(w http.ResponseWriter, tasks []models.Task, err error) {
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
func getTaskId(r *http.Request) (uint, error) {
	idString := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.ParseUint(idString, 10, 64) //конверт в uint
	if err != nil {
		return 0, appErrors.ErrInvalidTaskID
	}

	return uint(id), nil
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter) {
	tasks, err := h.service.GetActiveTasks()
	h.writeTasks(w, tasks, err)
}

func (h *TaskHandler) GetArchivedTasks(w http.ResponseWriter) {
	tasks, err := h.service.GetArchivedTasks()
	h.writeTasks(w, tasks, err)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, appErrors.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateTask(&task)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrTaskTextEmpty),
			errors.Is(err, appErrors.ErrTaskTextTooLong):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, appErrors.ErrInternalServerError.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println(err)
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(id)
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

func (h *TaskHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.CompleteTask(id)
	if err != nil {
		if errors.Is(err, appErrors.ErrTaskNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		log.Println(err)
		http.Error(w, appErrors.ErrInternalServerError.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println(err)
	}
}
