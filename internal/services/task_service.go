package services

import (
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
	"errors"
)

type TaskService struct {
	repository repositories.ITaskRepository
}

func NewTaskService(repository repositories.ITaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (service *TaskService) CreateTask(task *models.Task) error {
	if task.Text == "" {
		return errors.New("task text is empty")
	}

	// []rune, т.к len читает байты, а не символы
	if len([]rune(task.Text)) > 1000 {
		return errors.New("task text cannot be longer than 1000 characters")
	}

	return service.repository.Create(task)
}
