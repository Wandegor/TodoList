package router

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/handlers"
	"net/http"
)

func SetupRouter(taskHandler *handlers.TaskHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks/archive", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, appErrors.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
			return
		}

		taskHandler.GetArchivedTasks(w)
	})

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTasks(w)
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		default:
			http.Error(w, appErrors.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		}

	})

	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)
		case http.MethodPatch:
			taskHandler.CompleteTask(w, r)
		default:
			http.Error(w, appErrors.ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		}

	})

	return mux
}
