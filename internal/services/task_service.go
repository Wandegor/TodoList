package services

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
	"strings"
)

type TaskService struct {
	repository repositories.ITaskRepository
}

func NewTaskService(repository repositories.ITaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (service *TaskService) GetTasks() ([]models.Task, error) {

	return service.repository.GetAll()
}

func (service *TaskService) CreateTask(task *models.Task) error {
	if strings.TrimSpace(task.Text) == "" {
		return appErrors.ErrTaskTextEmpty
	}

	// []rune, т.к len читает байты, а не символы
	if len([]rune(task.Text)) > 1000 {
		return appErrors.ErrTaskTextTooLong
	}

	return service.repository.Create(task)
}

func (service *TaskService) DeleteTask(id uint) error {

	task, err := service.repository.GetByID(id)
	if err != nil {
		return err
	}

	return service.repository.Delete(&task)
}
